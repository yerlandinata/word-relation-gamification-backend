ALTER TABLE word_pair 
    ADD COLUMN active_status INTEGER;

UPDATE word_pair SET active_status = 1 WHERE active_status IS NULL;

ALTER TABLE word_pair
    ALTER COLUMN active_status SET NOT NULL;
    