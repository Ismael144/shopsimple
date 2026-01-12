.PHONY: buf run clean

buf: 
	buf generate buf.build/shopsimple/shopsimple
run: 
	go run cmd/grpc/main.go
clean: 
	rm -r gen