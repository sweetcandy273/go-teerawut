CREATE TABLE
     customers (
          id SERIAL PRIMARY KEY UNIQUE,
          created_at TIMESTAMPTZ DEFAULT now () NOT NULL,
          updated_at TIMESTAMPTZ DEFAULT now () NOT NULL,
          deleted_at TIMESTAMPTZ,
          name TEXT NOT NULL,
          surname TEXT NOT NULL,
          nickname TEXT,
          telephone_number TEXT,
          phone_number TEXT,
          detail TEXT,
          created_by_user_id INTEGER,
          updated_by_user_id INTEGER,
          deleted_by_user_id INTEGER
     );

-- เพิ่ม Index สำหรับ soft delete
CREATE INDEX idx_customers_deleted_at ON customers (deleted_at);

-- เพิ่ม Index สำหรับฟิลด์การค้นหาทั่วไป
CREATE INDEX idx_customers_name ON customers (name);

CREATE INDEX idx_customers_surname ON customers (surname);

CREATE INDEX idx_customers_nickname ON customers (nickname);

CREATE INDEX idx_customers_telephone_number ON customers (telephone_number);

CREATE INDEX idx_customers_phone_number ON customers (phone_number);

CREATE INDEX idx_customers_detail ON customers (detail);

ALTER TABLE customers ADD CONSTRAINT fk_customers_created_by FOREIGN KEY (created_by_user_id) REFERENCES users (id),
ADD CONSTRAINT fk_customers_updated_by FOREIGN KEY (updated_by_user_id) REFERENCES users (id),
ADD CONSTRAINT fk_customers_deleted_by FOREIGN KEY (deleted_by_user_id) REFERENCES users (id);