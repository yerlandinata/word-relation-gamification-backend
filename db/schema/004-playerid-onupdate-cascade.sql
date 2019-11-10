ALTER TABLE annotation
    DROP CONSTRAINT annotation_player_id_fkey,
    ADD CONSTRAINT annotation_player_id_fkey
        FOREIGN KEY (player_id) REFERENCES player(id) ON UPDATE CASCADE;
