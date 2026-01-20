import logging 

class RequestLoggerAdapter(logging.LoggerAdapter): 
    def process(self, msg, kwargs): 
        request_id = self.extra.get("request_id", "-")
        trace_id = self.extra.get("trace_id", "-")
        return msg, {"extra": {"request_id": request_id, "trace_id": trace_id}}

def get_request_id(context): 
    md = context.invocation_metadata()
    request_id = None 
    for key, value in md: 
        if key.lower() == "x-request-id":
            request_id = value 
    #  Fallback: generate if missing 
    if request_id is None: 
        import uuid 
        request_id = str(uuid.uuid4())
    return request_id