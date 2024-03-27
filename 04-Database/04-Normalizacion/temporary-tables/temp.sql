use movies_db;

-- Raw
CREATE TEMPORARY TABLE `temporary_table` (
	`id` int not null auto_increment,
    `name` varchar(250) not null,
    primary key (`id`)
);
INSERT INTO `temporary_table` (`id`, `name`) VALUES
(1, "titanic"),
(2, "buscano a nemo");

SELECT * FROM temporary_table;

-- TT genres
select * from genres;

CREATE TEMPORARY TABLE `genres_temp` LIKE genres;

INSERT INTO genres_temp
(SELECT * from genres);

SELECT * from genres_temp;