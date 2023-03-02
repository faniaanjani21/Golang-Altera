-- table product_types
-- pk id int(11)
--     name barchar(255)
--     create_at timestamp
--     update_at timestamp

CREATE TABLE product_types (
  id INT(11) PRIMARY KEY,
  name VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);


-- table operators
-- pk id int(11)
--     name barchar(255)
--     create_at timestamp
--     update_at timestamp
CREATE TABLE operators (
id int(11) NOT NULL AUTO_INCREMENT,
name varchar(255) NOT NULL,
created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
updated_at timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- table products_description
-- pk id int(11)
--     description text
--     create_at timestamp
--     update_at timestamp
CREATE TABLE products_description (
  id INT(11) PRIMARY KEY,
  description TEXT,
  create_at TIMESTAMP,
  update_at TIMESTAMP
);



-- table payment_methods
-- pk id int(11)
--     name varchar(255)
--     status smallint
--     create_at timestamp
--     update_at timestamp

CREATE TABLE payment_methods (
id int(11) NOT NULL,
name varchar(255) NOT NULL,
status smallint NOT NULL,
create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(),
update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP()
);

-- table users
-- pk id int(11)
--     status smallint
--     dob date
--     gender char(1)
--     create_at timestamp
--     update_at timestamp
CREATE TABLE users (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  status SMALLINT NOT NULL,
  dob DATE NOT NULL,
  gender CHAR(1) NOT NULL,
  create_at TIMESTAMP NOT NULL,
  updata_at TIMESTAMP NOT NULL
);


-- table transactions
-- pk id int(11)
-- fk user_id int(11)
-- fk payment_methods int(11)
-- status varchar(10)
-- total_qty int(11)
-- total_price numberic(25,2)
-- create_at timestamp
-- update_at timestamp

CREATE TABLE transactions (
 id INT(11) PRIMARY KEY, 
 user_id INT(11), 
 payment_methods INT(11), 
 status VARCHAR(10), 
 total_qty INT(11),
 total_price NUMERIC(25,2),
 create_at TIMESTAMP, 
 update_at TIMESTAMP
);


-- transactions_details
-- pk transactions_id int(11)
-- pk product_id int(11)
-- status varchar(10)
-- qty int(11)
-- price numberic(25,2)
-- create_at timestamp
-- update_at timestamp
CREATE TABLE transactions_details (
transactions_id int(11) NOT NULL,
product_id int(11) NOT NULL,
status varchar(10) NOT NULL,
qty int(11) NOT NULL,
price numeric(25,2) NOT NULL,
created_at timestamp NOT NULL,
update_at timestamp NOT NULL,
PRIMARY KEY (transactions_id, product_id)
);


-- table products
-- pk id int(11)
-- fk product_types_id int(11)
-- fk operators_id int(11)
-- code varchar(255)
-- name varchar(255)
-- status smallint
-- create_at timestamp
-- update_at timestamp

CREATE TABLE products (
id int(11) NOT NULL AUTO_INCREMENT,
product_types_id int(11) DEFAULT NULL,
operators_id int(11) DEFAULT NULL,
code varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
name varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
status smallint(6) DEFAULT NULL,
create_at timestamp NULL DEFAULT NULL,
update_at timestamp NULL DEFAULT NULL,
PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- tugas tugas nya
-- 1. Insert
--     1. Insert 5 operators pada table operators.
INSERT INTO operators (pk_id, name, create_at, update_at) 
VALUES 
(1, 'John Doe', current_timestamp(), current_timestamp()),
(2, 'Jane Doe', current_timestamp(), current_timestamp()),
(3, 'Bill Smith', current_timestamp(), current_timestamp()),
(4, 'Sarah Jones', current_timestamp(), current_timestamp()),
(5, 'Jill Johnson', current_timestamp(), current_timestamp());
--     2. Insert 3 product type.
INSERT INTO product_types(name, created_at, updated_at) VALUES
('Product Type 1', NOW(), NOW()),
('Product Type 2', NOW(), NOW()),
('Product Type 3', NOW(), NOW());
--     3. Insert 2 product dengan product type id = 1, dan operators id = 3. 
INSERT INTO product_types (id, name, create_at, update_at) 
VALUES (1, 'Product Name 1', NOW(), NOW());
INSERT INTO product_types (id, name, create_at, update_at) 
VALUES (3, 'Product Name 2', NOW(), NOW());
--     4. Insert 3 product dengan product type id = 2, dan operators id = 1.
INSERT INTO product_types (id, name, create_at, update_at)
VALUES (2, 'Product 1', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP()),
(2, 'Product 2', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP()),
(2, 'Product 3', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP());
--     5. Insert 3 product dengan product type id = 3, dan operators id = 4.
INSERT INTO product_types (name, create_at, update_at) VALUES ('Product Type 3', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP());

INSERT INTO products (type_id, operator_id) VALUES (3, 4);

INSERT INTO products (type_id, operator_id) VALUES (3, 4);

INSERT INTO products (type_id, operator_id) VALUES (3, 4);
--     6. Insert product description pada setiap product.
INSERT INTO products_description (id, description, create_at, update_at) 
VALUES (product_id, 'Deskripsi produk', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

--     7. Insert 3 payment methods.
INSERT INTO payment_methods (name, status, created_at, updated_at) 
VALUES 
('Visa', 0, CURDATE(), CURDATE()), 
('Mastercard', 0, CURDATE(), CURDATE()), 
('American Express', 0, CURDATE(), CURDATE());

--     8. Insert 5 user pada tabel user.
INSERT INTO users (status, dob, gender, create_at, update_at) VALUES (0, "YYYY-MM-DD", "M", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO users (status, dob, gender, create_at, update_at) VALUES (1, "YYYY-MM-DD", "F", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO users (status, dob, gender, create_at, update_at) VALUES (2, "YYYY-MM-DD", "O", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO users (status, dob, gender, create_at, update_at) VALUES (3, "YYYY-MM-DD", "M", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO users (status, dob, gender, create_at, update_at) VALUES (4, "YYYY-MM-DD", "F", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
--     9. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
INSERT INTO Transactions(user_id, payment_methods, status, total_qty, total_price, create_at, update_at) 
VALUES (1, 1, 'pending', 5, 10.00, NOW(), NOW()), 
(2, 2, 'completed', 3, 50.00, NOW(), NOW()), 
(3, 3, 'cancelled', 2, 20.00, NOW(), NOW());

--     10. Insert 3 product di masing-masing transaksi.
INSERT INTO products(id, product_types_id, operators_id, code, name, status, create_at, update_at)
VALUES(1, 123, 4, "ABC123", "Product 1", 1, NOW(), NOW()),
      (2, 234, 5, "XYZ456", "Product 2", 1, NOW(), NOW()),
      (3, 345, 6, "QRS789", "Product 3", 1, NOW(), NOW());

-- 2. Select
--     1. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT nama FROM customers_table WHERE gender = 'L' or gender = 'M';
db.customers.find( { gender: {$in: ["L", "M"]}}, {name: 1, _id: 0} );

--     2. Tampilkan product dengan id = 3.
SELECT * FROM products WHERE id = 3;

--     3. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
SELECT *  
FROM customers
WHERE created_at BETWEEN DATE_SUB(CURDATE(),INTERVAL 7 DAY) AND NOW() 
AND name LIKE '%a%'

--     4. Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(*) 
FROM user AS u
WHERE u.gender = 'Perempuan';

--     5. Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT (*) 
FROM users 
WHERE gender = 'Perempuan';

--     6. Tampilkan 5 data pada data product
SELECT * FROM Product LIMIT 5;

-- 3. Update
--     1. Ubah data product id 1 dengan nama ‘product dummy’.
UPDATE products SET name = 'product dummy' WHERE id = 1;
--     2. Update qty = 3 pada transaction detail dengan product id = 1.
UPDATE transaction_details 
SET qty = 3 
WHERE product_id = 1;

-- 4. Delete
--     1. Delete data pada tabel product dengan id = 1.
DELETE FROM product WHERE id = 1; 

--     2. Delete pada pada tabel product dengan product type id 1.
DELETE FROM product WHERE product_type_id = 1;
-- ### Join, Union, Sub query, Function

-- 1. Gabungkan data transaksi dari user id 1 dan user id 2.
SELECT 
'user_id_1' AS UserID, 
column_name_1,
column_name_2
FROM table_name_1
WHERE user_id = 1
UNION ALL
SELECT 
'user_id_2', 
column_name_1,
column_name_2
FROM table_name_2
WHERE user_id = 2

-- 2. Tampilkan jumlah harga transaksi user id 1.
SELECT SUM(price) 
FROM transactions 
WHERE user_id = 1;

-- 3. Tampilkan total transaksi dengan product type 2.
SELECT COUNT(*) AS total_transactions
FROM transactions
WHERE product_type = 2;

-- 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan.
SELECT product.field_name, product_type.field_name 
FROM product 
JOIN product_type ON product.product_type_id = product_type.product_type_id

-- 5. Tampilkan semua field table transaction, field name table product dan field name table user.
SELECT t.*, p.fieldName AS productFieldName, u.fieldName AS userFieldName
FROM transaction t
JOIN product p ON t.productID = p.productID
JOIN user u on t.userID = u.userID

-- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
DELETE FROM transactions  
WHERE transaction_id = <transaction_id> 
AND DELETE FROM transaction_details 
WHERE transaction_id = <transaction_id> 

-- 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.
CREATE FUNCTION UpdateTotalQty() 
RETURNS VOID 
AS
BEGIN 
  UPDATE Total_Qty SET qty = 
  (SELECT SUM(qty) 
  FROM TransactionDetails
  WHERE transaction_id NOT IN (SELECT transaction_id FROM DeletedTransaction))
END;

-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.
SELECT product.* 
FROM product 
WHERE product.id NOT IN (
    SELECT transaction_details.product_id 
    FROM transaction_details
)
