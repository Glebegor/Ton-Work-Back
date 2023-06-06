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
CREATE TABLE users_to_works (
    id serial not null unique,
    id_user int references users(id) on delete cascade not null,
    id_works int references works(id) on delete cascade not null
);