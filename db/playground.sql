SELECT 
    wp.id,
    wp.word_1,
    wp.word_2,
    wp.example_sentence,
    COUNT(*)
FROM word_pair wp
WHERE
    wp.id NOT IN (
        SELECT wp_id FROM gold_standard
    ) AND
    wp.id NOT IN (
        SELECT wp_id FROM annotation WHERE player_id=81290624825
    )
GROUP BY wp.id
ORDER BY COUNT(wp.id) DESC
;

SELECT
    p1.id,
    p1.full_name,
    p1.ranking 
FROM (
    SELECT 
        p2.id,
        p2.full_name,
        RANK() OVER (ORDER BY p2.score DESC) ranking
    FROM player p2 
) p1
WHERE p1.id=081290624825
;


SELECT
    wp.id,
    wp.word_1,
    wp.word_2,
    COUNT(a.wp_id)
FROM word_pair wp
LEFT JOIN annotation a on wp.id=a.wp_id
WHERE
    wp.id NOT IN (
        SELECT wp_id FROM annotation WHERE player_id=98086
    ) AND wp.id NOT IN (
        SELECT wp_id FROM gold_standard
    )
GROUP BY a.wp_id, wp.id
HAVING COUNT(a.wp_id) <= 2
ORDER BY COUNT(a.wp_id) DESC;

SELECT
        wp.id,
        wp.word_1,
        wp.word_2,
        COUNT(a.wp_id)
FROM word_pair wp
LEFT JOIN (SELECT wp_id FROM annotation WHERE is_valid=true) a on a.wp_id=wp.id

WHERE

wp.active_status=1
AND wp.id NOT IN (
        SELECT wp_id FROM annotation WHERE player_id=240743
)

AND wp.id NOT IN (
        SELECT wp_id FROM gold_standard
)

GROUP BY a.wp_id, wp.id

        HAVING COUNT(a.wp_id) < 7

ORDER BY COUNT(a.wp_id) DESC, wp.pmi, wp.word_1_freq DESC
LIMIT 20

insert into public.annotation (wp_id, wrt_id, player_id, player_time_ms, created_at, is_valid) VALUES (3203, 1, 11780, 0, CURRENT_TIMESTAMP, false);
insert into public.annotation (wp_id, wrt_id, player_id, player_time_ms, created_at, is_valid) VALUES (1870, 1, 11780, 0, CURRENT_TIMESTAMP, true);