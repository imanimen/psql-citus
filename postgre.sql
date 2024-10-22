-- HOW TO CREATE TABLE IN POSTGRESQL
create table users (
  id SERIAL primary key,
  name varchar(100),
  email varchar(100) unique,
  age int,
  created_at timestamp default current_timestamp
)

-- HOW TO INSERT IN POSTGRESQL
insert into users (name, email, age)
values
  ('iman', 'iman@gmail.com', 22),
  ('farid', 'farid@gmail.com', 25)

-- HOW TO QUERY DATA

select * from users;
select * from users where age > 23

-- HOW TO UPDATE DATA
update users 
set age = 18
where name = 'iman'

-- HOW TO DELETE DATA 
delete from users where name = 'iman'

-- ADVANCED
create table json_example (
  id serial primary key,
  data JSON
)


insert into json_example (data)
values ('{"name": "ali", "age": 28}')

-- QUERY JSON IN POSTGRESQL
select data->>'name' as name from json_example
FULL TEXT search
create table articles (
  id serial primary key,
  title text,
  content text
)

INSERT INTO articles (title, content) VALUES
('Introduction to PostgreSQL', 'PostgreSQL is a powerful, open-source object-relational database system.'),
('Full-Text Search in PostgreSQL', 'PostgreSQL supports full-text search, allowing powerful searching of text data.'),
('Getting Started with SQL', 'SQL is a standard language for managing relational databases.');

select *
from articles
where to_tsvector('english', content) @@ to_tsquery('english', 'PostgreSQL')

-- full text search on multiple columns (title and content)
select *
from articles
where to_tsvector('english', title || ' ' || content) @@ to_tsquery('english', 'PostgreSQL')
