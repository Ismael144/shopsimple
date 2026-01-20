import grpc 
import logging 
from concurrent import futures 
import os 
from adapters import jaegertracing
from dotenv import load_dotenv

# Grpc Imports 
from google.protobuf import empty_pb2
from shopsimple.currency.v1 import currency_pb2, currency_pb2_grpc
from currency_converter import convert_money, batch_convert_money, get_supported_currencies
from opentelemetry.sdk.trace import TracerProvider
from adapters import requestid, tracing
from opentelemetry import trace
from grpc_reflection.v1alpha import reflection


# Load environment variables 
load_dotenv()

# Init jaegar 
jaegertracing.InitJaegar(os.getenv("JAEGER_HOST"), os.getenv("JAEGER_PORT"))

# Get a tracer
tracer = trace.get_tracer(__name__) 

if not isinstance(trace.get_tracer_provider(), TracerProvider): 
    trace.set_tracer_provider(TracerProvider())

# Setup logging 
logging.basicConfig(
    level=logging.INFO, 
    format="%(asctime)s | %(levelname)s | %(request_id)s | %(trace_id)s | %(message)s"
)

class CurrencyService(currency_pb2_grpc.CurrencyServiceServicer): 
    def GetSupportedCurrencies(self, req: empty_pb2.Empty, context): 
        # Get propagated request id and trace id 
        request_id = requestid.get_request_id(context)
        trace_id = tracing.get_trace_id()
        # Init logger
        logger = requestid.RequestLoggerAdapter(logging.getLogger(__name__), {"request_id": request_id, "trace-id": trace_id})
        logger.info(f"Convert request received for supported currencies...")
        currencies = []
        with tracer.start_as_current_span("GetSupportedCurrencies"):
            currencies = get_supported_currencies()
        return currency_pb2.SupportedCurrenciesResponse(currencies=currencies)

    def Convert(self, req: currency_pb2.CurrencyConversionRequest, context):
        # Get propagated request id and trace id 
        request_id = requestid.get_request_id(context)
        trace_id = tracing.get_trace_id()
        # Init logger
        logger = requestid.RequestLoggerAdapter(logging.getLogger(__name__), {"request_id": request_id, "trace-id": trace_id})
        from_money_proto = {
            "currency_code": req.from_money.currency_code,
            "units": req.from_money.units,
            "nanos": req.from_money.nanos
        }
        try: 
            result = {}
            with tracer.start_as_current_span("Convert"):
                result = convert_money(from_money_proto, req.to_code)
            logger.info(f"Conversion successful: {result['units']}.{result['nanos']} {result['currency_code']}")
            return currency_pb2.shopsimple_dot_common_dot_v1_dot_money__pb2.Money(
                currency_code=result["currency_code"], 
                units=result["units"], 
                nanos=result["nanos"]
            )
        except ValueError as e: 
            logger.error(f"Conversion failed: {e}")
            context.set_details(str(e))
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            return currency_pb2.shopsimple_dot_common_dot_v1_dot_money__pb2.Money()

    def BatchConvert(self, req: currency_pb2.BatchCurrencyConversionRequest, context): 
        # Get propagated request id and trace id 
        request_id = requestid.get_request_id(context)
        trace_id = tracing.get_trace_id()
        # Init logger
        logger = requestid.RequestLoggerAdapter(logging.getLogger(__name__), {"request_id": request_id, "trace_id": trace_id})
        try:
            from_money_list = [
                {"currency_code": c.currency_code, "units": c.units, "nanos": c.nanos} for c in req.from_money
            ]
            results = batch_convert_money(from_money_list, req.to_code)
            return currency_pb2.BatchCurrencyConversionResponse(
                results=[
                        currency_pb2.shopsimple_dot_common_dot_v1_dot_money__pb2.Money(
                            currency_code=r["currency_code"], 
                            units=r["units"], 
                            nanos=r["nanos"],
                        )
                        for r in results 
                ]
            )
        except Exception as e:
            logger.error("An error occured: ", e)

def serve(): 
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    currency_pb2_grpc.add_CurrencyServiceServicer_to_server(
        CurrencyService(), 
        server
    )
    SERVICE_NAMES = (
        currency_pb2.DESCRIPTOR.services_by_name["CurrencyService"].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)
    server.add_insecure_port(f"[::]:{os.getenv("CURRENCYSERVICE_PORT")}")
    print(f"CurrencyService gRPC server running on port {os.getenv("CURRENCYSERVICE_PORT")}...")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__": 
    serve()