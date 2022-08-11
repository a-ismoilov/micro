create table meals(
    id serial not null primary key,
    name varchar(50) not null,
    price int
);

create table users(
    id serial not null primary key
);

create table orders(
    id serial not null primary key ,
    user_id int not null references users(id),
    opened_at timestamp default current_timestamp,
    closed_at timestamp default null,
    price int default 0
);