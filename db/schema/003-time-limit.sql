ALTER TABLE player 
    ADD COLUMN elapsed INTEGER;

UPDATE player SET elapsed = 0 WHERE elapsed IS NULL;

ALTER TABLE player
    ALTER COLUMN elapsed SET NOT NULL;
    