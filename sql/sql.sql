CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    username varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    userPassword varchar(20) not null unique,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;