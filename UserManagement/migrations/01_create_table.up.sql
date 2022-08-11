create table users(
    id serial not null primary key ,
    username varchar(32) not null ,
    budget int
);

create table admins(
    id serial not null primary key ,
    password varchar(32) not null
);