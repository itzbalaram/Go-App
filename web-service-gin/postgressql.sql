CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR ( 50 ) NOT NULL --this is name of the Author,
    category VARCHAR ( 50 ) UNIQUE NOT NULL,
    price VARCHAR ( 255 ) NOT NULL
);


SELECT id,name,category,price FROM products LIMIT 5 OFFSET 1;

-- UPDATE products 
-- SET name='nae56',category='tle56', price=66.7667
-- WHERE id=8;

SELECT * FROM products;

SELECT * FROM products WHERE id=8;

-- DELETE FROM products WHERE id=1;

INSERT INTO products(name,category,price) VALUES('name3','category3','price3') ON CONFLICT DO NOTHING;