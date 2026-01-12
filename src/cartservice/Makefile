.PHONY: buf run clean

buf: 
	buf generate buf.build/shopsimple/shopsimple
run: 
	go run cmd/grpc/main.go
clean: 
	rm -r gen
# Do some testing 
additem: 
	grpcurl -plaintext -d '{"product_id": "0510cc40-5612-471a-be92-28467aa46b55", "user_id": "someid", "quantity": 3}' localhost:50052 cart.v1.CartService/AddItem
cart: 
	grpcurl -plaintext -d '{"user_id": "someid"}' localhost:50052 cart.v1.CartService/GetCart
clear: 
	grpcurl -plaintext -d '{"user_id": "someid"}' localhost:50052 cart.v1.CartService/Clear
deduct: 
	grpcurl -plaintext -d '{"product_id": "0510cc40-5612-471a-be92-28467aa46b55", "user_id": "someid", "quantity": 1}' localhost:50052 cart.v1.CartService/DeductFromCart
