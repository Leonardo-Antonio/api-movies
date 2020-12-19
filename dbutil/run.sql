drop database if exists db_movies;
create database if not exists db_movies;
use db_movies;

select movies.id, movies.name, movies.stars, movies.state,
       concat(authors.name, " ", authors.last_name)
from movies
left join authors on movies.id = authors.movies_id;

select * from movies;
select * from categories;
select * from authors;

select * from authors
left join movies
on ( authors.movies_id = movies.id )
WHERE movies.categories_id = 1