DROP TABLE lists_items;

DROP TABLE users_list;

DROP TABLE todo_items;

DROP TABLE todo_lists;

DROP TABLE users;






-- migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down