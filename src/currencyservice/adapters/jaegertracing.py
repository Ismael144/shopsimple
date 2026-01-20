from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.exporter.jaeger.thrift import JaegerExporter

def InitJaegar(host: str | None, port: str | None): 
    host = "localhost" if host is None else host
    port = "6831" if port is None else port
    # Create a tracer provider
    trace.set_tracer_provider(TracerProvider())
    tracer = trace.get_tracer(__name__)

    # Configure Jaeger exporter
    jaeger_exporter = JaegerExporter(
        agent_host_name=host,  # your Jaeger agent host
        agent_port=port,              # default Jaeger UDP port
    )

    # Batch processor for sending spans asynchronously
    span_processor = BatchSpanProcessor(jaeger_exporter)
    trace.get_tracer_provider().add_span_processor(span_processor)

    print(f"OpenTelemetry configured to export traces to Jaeger at {host}:{port}")
