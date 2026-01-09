.PHONY: gen_proto 

gen_proto: 
		buf lint 
		buf generate
view_tables_in_db:
		docker exec -it products-db psql -U products_db_user -d productssvc -c "DROP TABLE schema_migrations;"
apply_migrations: 
		export DATABASE_URL=postgresql://products_db_user:products_db_password@localhost:5432/productssvc?sslmode=disable
		migrate -path ./migrations/ -database $DATABASE_URL up
