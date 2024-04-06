CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    date_created TIMESTAMP,
    last_updated TIMESTAMP
);

insert into users(first_name, last_name, email, date_created, last_updated)
values('John', 'Doe', 'john@doe.com', now(), now());