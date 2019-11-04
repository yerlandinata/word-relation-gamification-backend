INSERT INTO public.player (id, birth_date, full_name, education_level, score, annotation_count) VALUES (8123456789, 12345678, 'lolol', 'Sedang S1', 0, 0);
INSERT INTO public.player (id, birth_date, full_name, education_level, score, annotation_count) VALUES (987654321, 87654321, 'lele', 'Sedang SMA', 0, 0);
INSERT INTO public.player (id, birth_date, full_name, education_level, score, annotation_count) VALUES (81290624825, 27061998, 'Yudhistira Erlandinata', 'Sedang S1', 0, 0);
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (1, 'mobil', 'kendaraan');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (2, 'sepeda', 'kendaraan');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (3, 'mikroskop', 'alat');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (4, 'binatang', 'hewan');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (5, 'kucing', 'hewan');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (6, 'anjing', 'mamalia');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (17, 'komputer', 'sendok');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (18, 'sendal', 'kucing');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (13, 'memecah belah', 'tindakan');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (7, 'anjing', 'binatang');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (8, 'upah', 'gaji');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (9, 'upah', 'bayaran');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (10, 'gaji', 'bayaran');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (11, 'adu domba', 'memecah belah');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (12, 'adu domba', 'tindakan');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (14, 'manusia', 'orang');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (15, 'manusia', 'makhluk');
INSERT INTO public.word_pair (id, word_1, word_2) VALUES (16, 'orang', 'makhluk');
INSERT INTO public.word_relation_type (id, short_desc, human_desc) VALUES (2, 'synonymy', '%a sama artinya dengan %b');
INSERT INTO public.word_relation_type (id, short_desc, human_desc) VALUES (1, 'hyponymy', '%a adalah sejenis %b');
INSERT INTO public.word_relation_type (id, short_desc, human_desc) VALUES (3, 'unrelated', '');
INSERT INTO public.word_relation_type (id, short_desc, human_desc) VALUES (4, 'notsure', '');
INSERT INTO public.annotation (wp_id, wrt_id, player_id, player_time_ms) VALUES (2, 1, 81290624825, 10);
INSERT INTO public.annotation (wp_id, wrt_id, player_id, player_time_ms) VALUES (3, 1, 81290624825, 10);
INSERT INTO public.annotation (wp_id, wrt_id, player_id, player_time_ms) VALUES (4, 1, 81290624825, 10);
INSERT INTO public.gold_standard (wp_id, wrt_id) VALUES (1, 1);
INSERT INTO public.gold_standard (wp_id, wrt_id) VALUES (5, 1);
INSERT INTO public.gold_standard (wp_id, wrt_id) VALUES (14, 2);
INSERT INTO public.gold_standard (wp_id, wrt_id) VALUES (8, 2);
INSERT INTO public.gold_standard (wp_id, wrt_id) VALUES (17, 3);
INSERT INTO public.gold_standard (wp_id, wrt_id) VALUES (18, 3);