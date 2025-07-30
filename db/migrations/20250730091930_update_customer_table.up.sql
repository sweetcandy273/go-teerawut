ALTER TABLE customers
RENAME COLUMN fname TO name;

ALTER TABLE customers
DROP COLUMN lname,
DROP COLUMN nickname;