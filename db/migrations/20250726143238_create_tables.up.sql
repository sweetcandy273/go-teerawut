CREATE TABLE
     IF NOT EXISTS users (
          id BIGSERIAL PRIMARY KEY,
          username VARCHAR,
          password VARCHAR,
          display_name VARCHAR
     );

CREATE TABLE
     IF NOT EXISTS customers (
          id BIGSERIAL PRIMARY KEY,
          created_at TIMESTAMP NOT NULL,
          updated_at TIMESTAMP NOT NULL,
          deleted_at TIMESTAMP DEFAULT NULL,
          fname VARCHAR,
          lname VARCHAR,
          nickname VARCHAR,
          phone_number VARCHAR,
          detail TEXT,
          created_by_user_id BIGINT NOT NULL REFERENCES users (id),
          updated_by_user_id BIGINT NOT NULL REFERENCES users (id),
          deleted_by_user_id BIGINT REFERENCES users (id)
     );

CREATE TABLE
     IF NOT EXISTS customer_addresses (
          id BIGSERIAL PRIMARY KEY,
          created_at TIMESTAMP NOT NULL,
          updated_at TIMESTAMP NOT NULL,
          deleted_at TIMESTAMP DEFAULT NULL,
          customer_id BIGINT NOT NULL REFERENCES customers (id),
          address VARCHAR,
          village VARCHAR,
          moo VARCHAR,
          soi VARCHAR,
          road VARCHAR,
          sub_district VARCHAR,
          district VARCHAR,
          province VARCHAR,
          zip_code INTEGER,
          telephone_number VARCHAR,
          detail TEXT,
          created_by_user_id BIGINT NOT NULL REFERENCES users (id),
          updated_by_user_id BIGINT NOT NULL REFERENCES users (id),
          deleted_by_user_id BIGINT REFERENCES users (id)
     );

CREATE TABLE
     IF NOT EXISTS config_constants (
          id BIGSERIAL PRIMARY KEY,
          created_at TIMESTAMP NOT NULL,
          updated_at TIMESTAMP NOT NULL,
          deleted_at TIMESTAMP,
          "group" VARCHAR NOT NULL,
          name_en VARCHAR,
          name_th VARCHAR,
          option VARCHAR,
          description VARCHAR,
          sort INTEGER,
          is_active SMALLINT NOT NULL
     );

CREATE TABLE
     IF NOT EXISTS customer_air_conditions (
          id BIGSERIAL PRIMARY KEY,
          created_at TIMESTAMP NOT NULL,
          updated_at TIMESTAMP NOT NULL,
          deleted_at TIMESTAMP DEFAULT NULL,
          customer_id BIGINT NOT NULL REFERENCES customers (id),
          air_condition_id BIGINT REFERENCES config_constants (id),
          btu_id BIGINT REFERENCES config_constants (id),
          room_name VARCHAR,
          from_us SMALLINT NOT NULL
     );

CREATE TABLE
     IF NOT EXISTS customer_service_lists (
          id BIGSERIAL PRIMARY KEY,
          created_at TIMESTAMP NOT NULL,
          updated_at TIMESTAMP NOT NULL,
          deleted_at TIMESTAMP DEFAULT NULL,
          customer_id BIGINT NOT NULL REFERENCES customers (id),
          date DATE NOT NULL,
          price NUMERIC(15, 2) NOT NULL,
          description VARCHAR,
          created_by_user_id BIGINT NOT NULL REFERENCES users (id),
          updated_by_user_id BIGINT NOT NULL REFERENCES users (id),
          deleted_by_user_id BIGINT REFERENCES users (id)
     );

CREATE TABLE
     IF NOT EXISTS customer_services (
          id BIGSERIAL PRIMARY KEY,
          customer_service_list_id BIGINT NOT NULL REFERENCES customer_service_lists (id),
          type_id BIGINT NOT NULL REFERENCES config_constants (id),
          description VARCHAR
     );

CREATE TABLE
     IF NOT EXISTS customer_service_items (
          customer_service_id BIGINT REFERENCES customer_services (id),
          item_id BIGINT REFERENCES config_constants (id)
     );