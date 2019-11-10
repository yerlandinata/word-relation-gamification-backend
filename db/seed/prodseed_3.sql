UPDATE public.gold_standard SET wrt_id=3 WHERE wp_id=3016;
INSERT INTO public.gold_standard (wp_id, wrt_id) VALUES (1796, 2);
UPDATE public.player SET score=0;
INSERT INTO public.player (id, full_name, birth_date, education_level, score, annotation_count, elapsed, game_level) VALUES 
    (1, 'Shiro', 0, '', 509, 0, 0, 10),
    (2, 'Finn', 0, '', 200, 0, 0, 6),
    (3, 'AAA', 0, '', 159, 0, 0, 4),
    (4, 'Jon', 0, '', 140, 0, 0, 4),
    (5, 'Bro', 0, '', 90, 0, 0, 3),
    (6, 'Kazuto', 0, '', 87, 0, 0, 3),
    (7, 'El', 0, '', 85, 0, 0, 3),
    (8, 'Doremi', 0, '', 82, 0, 0, 3),
    (9, 'Capucino', 0, '', 81, 0, 0, 3),
    (10, 'Kapten', 0, '', 79, 0, 0, 3);
