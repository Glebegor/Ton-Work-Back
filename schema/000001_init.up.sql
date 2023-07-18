CREATE TABLE users (
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    email varchar(255) not null,
    telefon varchar(255),
    position varchar(255),
    description varchar(2000),
    subscribe varchar not null,
    companies varchar,
    Name varchar,
    Surname varchar,
    id serial not null unique
);
CREATE TABLE works (
    title varchar(255),
	description varchar(255),
	text varchar(2000),
	tags varchar(2000),
	technologies varchar(2000),
	company varchar(2000),
	price int,
	experienceLevel varchar(200),
	type_of_job varchar(255),
	invites int,
	rating int,
	id serial not null unique
);
CREATE TABLE posts (
    title varchar(255),
    description varchar(255),
    text varchar(2000),
    tags varchar(2000),
    rating int,
    id serial not null unique
);
CREATE TABLE users_to_works (
    id serial not null unique,
    id_user int references users(id) on delete cascade not null,
    id_works int references works(id) on delete cascade not null
);
CREATE TABLE users_to_posts (
    id serial not null unique,
    id_user int references users(id) on delete cascade not null,
    id_posts int references posts(id) on delete cascade not null
);
CREATE TABLE user_sub (
    time_in_hours_to_end int,
    id_user int references users(id) on delete cascade not null,
    id serial not null unique
);

-- CREATE TABLE users (
--     id INT IDENTITY(1,1) PRIMARY KEY,
--     username VARCHAR(255) NOT NULL UNIQUE,
--     password_hash VARCHAR(255) NOT NULL,
--     email VARCHAR(255) NOT NULL,
--     telefon VARCHAR(255),
--     position VARCHAR(255),
--     description VARCHAR(2000),
--     subscribe VARCHAR(255) NOT NULL,
--     companies VARCHAR(255),
--     Name VARCHAR(255),
--     Surname VARCHAR(255)
-- );

-- CREATE TABLE works (
--     id INT IDENTITY(1,1) PRIMARY KEY,
--     title VARCHAR(255),
--     description VARCHAR(255),
--     text VARCHAR(2000),
--     tags VARCHAR(2000),
--     technologies VARCHAR(2000),
--     company VARCHAR(2000),
--     price INT,
--     experienceLevel VARCHAR(200),
--     type_of_job VARCHAR(255),
--     invites INT,
--     rating INT
-- );

-- CREATE TABLE posts (
--     id INT IDENTITY(1,1) PRIMARY KEY,
--     title VARCHAR(255),
--     description VARCHAR(255),
--     text VARCHAR(2000),
--     tags VARCHAR(2000),
--     rating INT
-- );

-- CREATE TABLE users_to_works (
--     id INT IDENTITY(1,1) PRIMARY KEY,
--     id_user INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
--     id_works INT REFERENCES works(id) ON DELETE CASCADE NOT NULL
-- );

-- CREATE TABLE users_to_posts (
--     id INT IDENTITY(1,1) PRIMARY KEY,
--     id_user INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
--     id_posts INT REFERENCES posts(id) ON DELETE CASCADE NOT NULL
-- );

-- CREATE TABLE user_sub (
--     id INT IDENTITY(1,1) PRIMARY KEY,
--     time_in_hours_to_end INT,
--     id_user INT REFERENCES users(id) ON DELETE CASCADE NOT NULL
-- );