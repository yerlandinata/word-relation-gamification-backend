ALTER TABLE player 
    ADD COLUMN game_level INTEGER;

UPDATE player SET game_level = 1 WHERE game_level IS NULL;

ALTER TABLE player
    ALTER COLUMN game_level SET NOT NULL;
    