-- Product categories table 
CREATE TABLE product_categories(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY, 
    name VARCHAR NOT NULL, 
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
); 

-- Products table
CREATE TABLE products (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY, 
    name VARCHAR NOT NULL, 
    description TEXT NULL, 
    unit_price BIGINT NOT NULL, 
    stock INT NOT NULL, 
    category_id FOREIGN KEY NOT NULL REFERENCES product_categories(id)
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
); 
