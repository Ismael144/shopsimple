# Design for a consistent inventory service

### Challenge(Scenario)
Let's say you have 2 users on your website or ecommerce application, they both move navigate around your application looking for products they may need, so they add the product the chose into their shopping carts respectively, after adding to cart, they head to checkout, where the service captures their addresses and credit card account details inorder to proceed with payment.

But, before payment, lets say the checkout service has to validate stock of the items in the cart but contacting/reaching out to inventory service to validate stock, in a situation where those 2 users at the same time added the same item in cart, but that product has a stock quantity of e.g.
10 but user 1 and user 2 both added that same item with quantities of 10, making it 20 total, but remember
the stock quantity is 10, since the cart service only validates the stock for each individual user using 
inventory service, the cart will not be able to catch 20 items that were taken, so when user 1, successfully 
makes payment, order is created and their products are shipped, and user 2 is just paying now and then the 
checkout service proceeds to validate stock finding out that the stock is 0, and there are 10 items in the 
cart, causing an error, but we wouldn't want a user to encounter such an error, or inconsistency, its so bad
So, what do we do to solve the problem.

### What I would do 
Lets say our inventory service is using postgres as database, or any other(doesn't matter now). 
According to me, I would add redis 