FROM envoyproxy/envoy:tools-dev-935971a4271f1f2aff11056f78f2dd309681311e

COPY envoy.yaml /etc/envoy/envoy.yaml 
COPY descriptor.pb /etc/envoy/descriptor.pb 

CMD ["envoy", "-c", "/etc/envoy/envoy.yaml", "--log-level", "info"]
