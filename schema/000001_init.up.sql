CREATE TABLE users (
    id INT IDENTITY(1,1) PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(2000) NOT NULL,
    email VARCHAR(255) NOT NULL,
    telefon VARCHAR(255),
    position VARCHAR(255),
    description VARCHAR(2000),
    subscribe VARCHAR(255) NOT NULL,
    companies VARCHAR(255),
    name VARCHAR(255),
    surname VARCHAR(255)
);

CREATE TABLE works (
    id INT IDENTITY(1,1) PRIMARY KEY,
    title VARCHAR(255),
    description VARCHAR(255),
    text VARCHAR(2000),
    tags VARCHAR(2000),
    technologies VARCHAR(2000),
    company VARCHAR(2000),
    price INT,
    experienceLevel VARCHAR(200),
    type_of_job VARCHAR(255),
    invites INT,
    rating INT
);

CREATE TABLE posts (
    id INT IDENTITY(1,1) PRIMARY KEY,
    title VARCHAR(255),
    description VARCHAR(255),
    text VARCHAR(2000),
    tags VARCHAR(2000),
    rating INT
);

CREATE TABLE users_to_works (
    id INT IDENTITY(1,1) PRIMARY KEY,
    id_user INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    id_works INT REFERENCES works(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE users_to_posts (
    id INT IDENTITY(1,1) PRIMARY KEY,
    id_user INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    id_posts INT REFERENCES posts(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE user_sub (
    id INT IDENTITY(1,1) PRIMARY KEY,
    time_in_hours_to_end INT,
    id_user INT REFERENCES users(id) ON DELETE CASCADE NOT NULL
);