USE movies_db;

-- 1. Mostrar todos los registros de la tabla de movies.
SELECT * FROM movies;

-- 2. Mostrar el nombre, apellido y rating de todos los actores.
SELECT ac.`first_name`, ac.`last_name`, ac.`rating` FROM actors as ac;


-- 3. Mostrar el título de todas las series y usar alias
-- para que tanto el nombre de la tabla como el campo estén en español.
SELECT se.title as 'titulo' FROM series as se;


-- Mostrar el nombre y apellido de los actores cuyo rating sea mayor a 7.5.
select a.first_name, a.last_name, a.rating from actors a
where a.rating > 7.5;

-- Mostrar el título, rating y los premios de las películas
-- con un rating mayor a 7.5 y
-- con más de dos premios.
SELECT m.title, m.rating, m.awards FROM movies as m
WHERE m.rating > 7.5 AND m.awards > 2;

-- Mostrar el título de las películas y el rating ordenadas por rating en forma ascendente.
SELECT m.title, m.rating FROM movies as m
ORDER BY m.rating;

-- Mostrar los títulos de las primeras tres películas en la base de datos.
SELECT m.title from movies as m
LIMIT 3;

-- Mostrar el top 5 de las películas con mayor rating.
SELECT m.* FROM movies as m
ORDER BY m.rating DESC
LIMIT 5;

-- Listar los primeros 10 actores.
SELECT a.* FROM actors as a
LIMIT 10;

-- Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.
SELECT m.title, m.rating FROM movies as m
WHERE m.title LIKE '%Toy Story%';

-- Mostrar a todos los actores cuyos nombres empiezan con Sam.
SELECT a.* FROM actors as a
WHERE a.first_name LIKE 'Sam%';

-- Mostrar el título de las películas que salieron entre el 2004 y 2008.
SELECT m.title from movies as m
-- WHERE m.release_date BETWEEN '2004-01-01 00:00:00' AND '2008-12-31 23:59:59';
-- WHERE m.release_date BETWEEN '2004-01-01 00:00:00' AND '2008-12-31 23:59:59';
WHERE YEAR(m.release_date) BETWEEN 2004 AND 2008;

-- Traer el título, rating y awards de las películas con el rating mayor a 3,
-- con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009.
-- Ordenar los resultados por rating de mayor a menor
SELECT m.title, m.rating, m.awards FROM movies as m
WHERE m.rating > 3 AND m.awards > 1 AND
m.release_date BETWEEN '1998-01-01' AND '2009-12-31'
ORDER BY m.rating DESC, m.awards DESC;
