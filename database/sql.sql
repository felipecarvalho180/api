CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

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
) ENGINE=INNODB;

CREATE TABLE posts(
  id int auto_increment primary key,
  title varchar(50) not null,
  content varchar(300) not null,

  author_id int not null,
  FOREIGN KEY (author_id)
  REFERENCES users(id)
  ON DELETE CASCADE,

  likes int default 0,
  createdAt timestamp default current_timestamp()
) ENGINE=INNODB;