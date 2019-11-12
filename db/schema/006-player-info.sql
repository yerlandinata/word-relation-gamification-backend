ALTER TABLE player 
    RENAME COLUMN full_name TO display_name;

ALTER TABLE player 
    ADD COLUMN full_name VARCHAR(64) NULL,
    ADD COLUMN onboarding_time_ms INTEGER NULL;
