CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS followers;

CREATE TABLE users(
  id int auto_increment primary key,
  name varchar(50) not null,
  nick varchar(50) not null unique,
  email varchar(50) not null unique,
  password varchar(100) not null,
  createdAt timestamp default current_timestamp()
 ) ENGINE=INNODB;

CREATE TABLE followers(
  userID int not null, 
  FOREIGN KEY(userID)
  REFERENCES users(id)
  ON DELETE CASCADE,

  followerID int not null, 
  FOREIGN KEY(followerID)
  REFERENCES users(id)
  ON DELETE CASCADE,

  primary key(userID, followerID)
)