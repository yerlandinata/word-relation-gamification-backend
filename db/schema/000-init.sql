CREATE TABLE word_pair(
   id serial PRIMARY KEY,
   word_1 VARCHAR (64) NOT NULL,
   word_2 VARCHAR (64) NOT NULL
);

CREATE TABLE word_relation_type(
    id serial PRIMARY KEY,
    short_desc VARCHAR(16) NOT NULL,
    human_desc VARCHAR(32) NOT NULL
);

CREATE TABLE gold_standard(
    wp_id INTEGER REFERENCES word_pair(id),
    wrt_id INTEGER REFERENCES word_relation_type(id),
    PRIMARY KEY (wp_id)
);

CREATE TABLE player(
    id BIGINT PRIMARY KEY,
    birth_date INTEGER NOT NULL,
    full_name VARCHAR(64) NOT NULL,
    education_level VARCHAR(64) NOT NULL,
    score INTEGER NOT NULL,
    annotation_count INTEGER NOT NULL
);

CREATE TABLE annotation(
    wp_id INTEGER REFERENCES word_pair(id),
    wrt_id INTEGER REFERENCES word_relation_type(id),
    player_id BIGINT REFERENCES player(id),
    player_time_ms INTEGER NOT NULL,
    PRIMARY KEY (wp_id, player_id)
);

CREATE INDEX annotation_player_id ON annotation (player_id);

INSERT INTO public.word_relation_type (id, short_desc, human_desc) VALUES (2, 'synonymy', '%a sama artinya dengan %b');
INSERT INTO public.word_relation_type (id, short_desc, human_desc) VALUES (1, 'hyponymy', '%a adalah sejenis %b');
INSERT INTO public.word_relation_type (id, short_desc, human_desc) VALUES (3, 'unrelated', '');
INSERT INTO public.word_relation_type (id, short_desc, human_desc) VALUES (4, 'notsure', '');

