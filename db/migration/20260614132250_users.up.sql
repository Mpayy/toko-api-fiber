create table if not exists users (
    id int auto_increment primary key,
    username varchar(255) not null,
    password varchar(255) not null,
    email varchar(255) not null unique,
    token varchar(255) null,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp on update current_timestamp
);