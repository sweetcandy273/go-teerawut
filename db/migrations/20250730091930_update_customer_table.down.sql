ALTER TABLE customers
RENAME COLUMN name TO fname;

ALTER TABLE customers
ADD COLUMN lname TEXT,
ADD COLUMN nickname TEXT;