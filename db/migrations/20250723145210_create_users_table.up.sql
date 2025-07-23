CREATE TABLE users (
    id SERIAL PRIMARY KEY UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    username TEXT NOT NULL,
    password TEXT NOT NULL
);

-- เพิ่ม Index ให้กับ deleted_at (ตาม gorm:"index")
CREATE INDEX idx_users_deleted_at ON users (deleted_at);
