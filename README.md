# Shopsimple Online Store
Shopsimple is a simple ecommerce online store built with intent of exploring the vastness of [microservice architecture](wikipedia.com)

## System Design 

### Functional Requirements 
1. Users can search and filter products by categories and prices
2. They can add and remove items from cart even without initially logging in
3. Users and login and register
4. They can checkout and make payments with credit card 
5. When done with checkout, they can receive an email of a shipping tracking id 

### Non-Functional Requirements 
1. **Scalable**: Can scale easily amongst high traffic spikes.<br/>
2. **Reliable**: Is predictable, does what the user expects it to do.<br/>
3. **Fault tolerance**: Can continue to operate even when some components have failed or are misbehaving.<br/>
4. **Partition tolerance**: Able to tolerate network partitions without users noticing.<br/>
5. **Availability**: Users won't face any downtime whilst using the application 

### Estimation & Constraints 
#### Users
Monthly number of users = **500M**<br/>
Monthly active number of users(MAU) = **250M**<br/>

#### RPS
Requests Per Second per day(MAU / 86400) = 250M / 86400 = **3000** RPS(Roughly) <br/>

#### Data
For **500M** Users, let's say record is 100 Bytes
<br/>
**500M** Users x **100** Bytes = 50GB(Roughly)  

# Services Schemas

## 1. Product Catalog Service<[ProductSvc]> (MongoDB)
**Products table**

- **id**: String[UUID]
- **name**: String 
- **description**: String 
- **rating**: Int 
- **unit_price**: Type {
        **CurrencyCode**: String, **Units**: Int64, **Nanos**: Int32,
    }
- **img_url**: String[URL]
- **stock**: Int
- **categories**: Array[String]
- **date_created**: Date

## 2. Cart Service<[CartSvc]> (Redis)
**Cart Item**
- **product_id**: > {ProductSvc}->(id)[UUID]
- **product_name**: String->{ProductSvc}
- **unit_price**: Type {
        **CurrencyCode**: String, **Units**: Int64, **Nanos**: Int32,
    }
- **quantity**: Int 

## 3. Auth Service<[AuthSvc]> (Postgres) 
**Users Table** 

- **id**: UUID Primary Key
- **username**: String 
- **email**: String
- **password**: String
- **role**: Enum {
        User, Admin
    }
- **is_active**: Boolean 
- **date_created**: Date

## 4. Checkout Service<[CheckoutSvc]> (Postgres)
**User Addresses Table [addresses]**

- **id**: UUID Primary Key 
- **user_id**: {AuthSvc}->(id)[UUID]
- **street_address**: String 
- **city**: String 
- **state**: String 
- **country**: String 
- **zipcode**: String 

**Orders Table**
- **id**: UUID Primary key 
- **user_id**: {AuthSvc}->(id)[UUID]
- **order_number**: String **[Auto-Generated]**
- **created_at**: Date

**Order Items Table**
- **id**: UUID Primary Key 
- **order_id**: **{Self:(Orders Table)}**->(id)[UUID]
- **cart_item**: **{CartSvc}->({CartSvc:(Cart Item)})**[JSONB]
- **status**: Enum {
        Shipped, Pending
    }

## 5. Shipping Service (Postgres)
- **id**: UUID Primary Key 
- **user_id**: {AuthSvc(Users Table)}->(id)[UUID]
- **address_id**: {CheckSvc(User Addresses Table)}->(id)[UUID]
- **tracking_id**: String **[Auto-Generated]**
- **status**: Enum {
    Received, Pending
}

## 6. Currency Service (Sqlite)
**Currency Conversion Rates Table(currency_conversions)**

- **id**: INT AUTO-INCREMENT Primary Key 
- **currency_code**: String[Format:**ISO-4217**]
- **rates**: Float