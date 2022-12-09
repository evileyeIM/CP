CREATE TABLE users
(
    id       serial       not null unique,
    email    varchar(255) not null,
    name     varchar(255) not null,
    password varchar(255) not null
);

CREATE TABLE notes
(
    userId          serial       not null unique,
    name            varchar(255) not null,
    date            varchar(255) not null,
    value           varchar(255)
);