-- Terdapat sebuah query pada SQL yaitu `SELECT * FROM users;`

-- Dengan tujuan yang sama, tuliskan dalam bentuk perintah:

-- SQL
SELECT * FROM users;
-- 1. Redis
GET * FROM users
-- 2. Neo4j
MATCH (n:User) RETURN n
-- 3. Cassandra
SELECT * FROM users_table