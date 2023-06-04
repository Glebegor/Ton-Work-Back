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
    Title varchar(255),
	Description varchar(255),
	Text varchar(2000),
	Tags varchar(2000),
	Technologies varchar(2000),
	Company varchar(2000),
	Price int,
	ExperienceLevel varchar(200),
	type_of_job varchar(255),
	invites int,
	Rating int,
	Id serial not null unique
);
CREATE TABLE users_to_works (
    id serial not null unique,
    id_user int references users(id) on delete cascade not null,
    id_works int references works(id) on delete cascade not null
);