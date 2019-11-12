ALTER TABLE player 
    ALTER COLUMN birth_date DROP NOT NULL;

ALTER TABLE player 
    RENAME COLUMN education_level TO campaign_source;
