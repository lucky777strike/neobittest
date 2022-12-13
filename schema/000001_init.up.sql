CREATE TABLE cams
(
    id SERIAL PRIMARY KEY,
    camname CHARACTER VARYING(30),
    lon double precision,
    lat double precision,
    addr CHARACTER VARYING(50) UNIQUE
);

CREATE TABLE tasks
(
    id SERIAL PRIMARY KEY,
    query CHARACTER VARYING(30),
    status CHARACTER VARYING(60),
    count smallint
);