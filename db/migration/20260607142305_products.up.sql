create table if not exists products(
    id int auto_increment primary key,
    name varchar(255) not null,
    price int not null,
    stock int not null,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp on update current_timestamp
);