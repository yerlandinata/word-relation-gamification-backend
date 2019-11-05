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
