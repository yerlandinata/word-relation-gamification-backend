SELECT
    wp.word_1,
    wp.word_2,
    wrt.short_desc
FROM word_pair wp
JOIN gold_standard g ON g.wp_id=wp.id
JOIN word_relation_type wrt ON wrt.id=g.wrt_id;
