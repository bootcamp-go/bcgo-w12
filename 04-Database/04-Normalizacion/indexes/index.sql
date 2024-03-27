use movies_db;

CREATE TEMPORARY TABLE `genres_temp` LIKE genres;

INSERT INTO genres_temp
(SELECT * from genres);

-- DQL
SELECT * FROM genres_temp;

-- doesn't do full table scan, primary key is already a unique key index
SELECT * from genres_temp WHERE id = 5;

-- ranking
SELECT * FROM genres_temp WHERE ranking = 5;

ALTER TABLE `genres_temp` ADD INDEX `idx_genres_temp_name` (`name`);
ALTER TABLE `genres_temp` ADD INDEX `idx_genres_unique` (`ranking`);
ALTER TABLE `genres_temp` DROP INDEX `idx_genres_temp_name`;
ALTER TABLE `genres_temp` DROP INDEX `idx_genres_unique`;

SELECT * FROM genres_temp WHERE `name` = "Comedia";

-- One indexed Other Not
SELECT * FROM genres_temp WHERE ranking = 6 AND `name` = "Suspenso";

SELECT * FROM genres_temp WHERE ranking = 6 OR `name` = "Comedia";

ALTER TABLE `genres_temp` ADD INDEX `idx_genres_temp_ranking_name` (`ranking`, `name`);
ALTER TABLE `genres_temp` DROP INDEX `idx_genres_temp_ranking_name`;

