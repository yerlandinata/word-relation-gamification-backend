ALTER TABLE annotation 
    ADD COLUMN is_valid BOOLEAN;

UPDATE annotation SET is_valid = true WHERE is_valid IS NULL;

ALTER TABLE annotation
    ALTER COLUMN is_valid SET NOT NULL;
    