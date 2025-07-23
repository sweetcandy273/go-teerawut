-- ลบ Index ที่เคยสร้างไว้
DROP INDEX IF EXISTS idx_customers_name;

DROP INDEX IF EXISTS idx_customers_surname;

DROP INDEX IF EXISTS idx_customers_nickname;

DROP INDEX IF EXISTS idx_customers_telephone_number;

DROP INDEX IF EXISTS idx_customers_phone_number;

DROP INDEX IF EXISTS idx_customers_detail;

DROP INDEX IF EXISTS idx_customers_deleted_at;

-- ลบ Foreign Keys (ถ้ามีเพิ่มไว้)
ALTER TABLE IF EXISTS customers
DROP CONSTRAINT IF EXISTS fk_customers_created_by,
DROP CONSTRAINT IF EXISTS fk_customers_updated_by,
DROP CONSTRAINT IF EXISTS fk_customers_deleted_by;

-- ลบ Table
DROP TABLE IF EXISTS customers;