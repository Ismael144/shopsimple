-- Product categories table 
CREATE TABLE IF NOT EXISTS product_categories(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY, 
    name VARCHAR NOT NULL, 
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
); 

-- Products table
CREATE TABLE IF NOT EXISTS products (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY, 
    name VARCHAR UNIQUE NOT NULL, 
    description TEXT NULL, 
    unit_price FLOAT NOT NULL, 
    image_url TEXT NULL,
    stock INT NOT NULL, 
    category_id UUID NOT NULL REFERENCES product_categories(id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
); 
