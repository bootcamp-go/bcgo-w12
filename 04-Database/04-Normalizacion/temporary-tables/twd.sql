SELECT ss.id as serie_id, ss.title as serie_title, se.id as season_id, se.title as season_title, ep.id as episode_id, ep.title as episode_title
FROM series ss
INNER JOIN seasons se ON ss.id = se.serie_id
INNER JOIN episodes ep ON se.id = ep.season_id
WHERE ss.title = "The Walking Dead";

CREATE TEMPORARY TABLE `twd_temp` as (
SELECT ss.id as serie_id, ss.title as serie_title, se.id as season_id, se.title as season_title, ep.id as episode_id, ep.title as episode_title
FROM series ss
INNER JOIN seasons se ON ss.id = se.serie_id
INNER JOIN episodes ep ON se.id = ep.season_id
WHERE ss.title = "The Walking Dead"
);

select * from twd_temp;