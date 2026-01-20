.PHONY: buf build run view_tables_in_db apply_migrations

buf: 
		buf generate buf.build/shopsimple/shopsimple
build: 
		docker compose build productservice 
		docker compose up productservice -d 
run: 
		go run cmd/grpc/main.go
exec_db:
		docker exec -it products-db psql -U products_db_user -d productssvc
apply_migrations: 
		export DATABASE_URL=postgresql://products_db_user:products_db_password@localhost:5432/productssvc?sslmode=disable
		migrate -path ./migrations/ -database $DATABASE_URL up
test_filter_service: 
		grpcurl -plaintext \
		-d '{
			"page": 1,
			"page_size": 20,
			"search_string": "modern",
			"categories": ["c1eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"],
			"price_ranges": {
			"min": 10.50,
			"max": 100000.00
			}
		}' \
		localhost:50051 \
		shopsimple.product.v1.ProductService/Filter
test_create_product: 
		grpcurl -plaintext \
		-d '{
			"name": "some product", 
			"description": "some description", 
			"unit_price": {
				"currency_code": "USD", 
				"units": 100, 
				"nanos": 32
			}, 
			"image_url": "https://images-na.ssl-images-amazon.com/images/G/01/AmazonExports/Events/2024/Stores-Gaming/FinalGraphics/Fuji_Gaming_store_Dashboard_card_1x_EN._SY304_CB564799420_.jpg", 
			"stock": 100,
			"rating": 5, 
			"categories": ["gaming", "pc", "rgb"]
			}' \
		localhost:50051 \
		shopsimple.product.v1.ProductService/Ctreate