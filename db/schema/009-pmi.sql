ALTER TABLE word_pair 
    ADD COLUMN pmi DOUBLE PRECISION;

UPDATE word_pair SET pmi = cast('0' as DOUBLE PRECISION) WHERE pmi IS NULL;

ALTER TABLE word_pair
    ALTER COLUMN pmi SET NOT NULL;
    