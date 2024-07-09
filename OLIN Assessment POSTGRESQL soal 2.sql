-- SELECT * FROM users

CREATE EXTENSION postgres_fdw
CREATE SERVER second_db_server
FOREIGN DATA WRAPPER postgres_fdw
OPTIONS (host 'localhost', dbname 'second', port '5432');

CREATE USER MAPPING FOR CURRENT_USER
SERVER second_db_server
OPTIONS (user 'postgres', password 'root');

IMPORT FOREIGN SCHEMA public
FROM SERVER second_db_server
INTO public;

SELECT u.name, o.amount, o.created_at
FROM users u
JOIN orders o ON u.id = o.user_id;