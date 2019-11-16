SELECT
    p.display_name,
    wp.word_1,
    wp.word_2,
    wrta.short_desc as ans,
    wrtg.short_desc as gold,
    a.player_time_ms
FROM annotation a
LEFT JOIN player p on p.id=a.player_id
LEFT JOIN word_pair wp on wp.id=a.wp_id
LEFT JOIN gold_standard g on g.wp_id=wp.id
LEFT JOIN word_relation_type wrta on wrta.id=a.wrt_id
LEFT JOIN word_relation_type wrtg on wrtg.id=g.wrt_id
WHERE player_id=8156168089
ORDER BY wp.word_2 DESC, wp.word_1 DESC;
