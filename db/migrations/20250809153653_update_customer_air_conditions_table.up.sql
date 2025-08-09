ALTER TABLE customer_air_conditions
ALTER COLUMN from_us
DROP DEFAULT,
ALTER COLUMN from_us TYPE boolean USING (
     CASE
          WHEN from_us IS NULL THEN NULL
          ELSE from_us <> 0
     END
),
ALTER COLUMN from_us
DROP NOT NULL;