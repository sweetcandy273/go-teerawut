ALTER TABLE customer_air_conditions
RENAME COLUMN air_condition_id TO air_brand_id;

ALTER TABLE customer_air_conditions
ADD COLUMN air_type_id BIGINT REFERENCES config_constants (id);