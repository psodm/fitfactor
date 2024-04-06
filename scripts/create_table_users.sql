CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS sex;

CREATE TYPE sex AS ENUM ('Male', 'Female');

CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    birthday DATE NOT NULL,
    height NUMERIC NOT NULL,
    sex sex NOT NULL,
    date_created TIMESTAMP DEFAULT now(),
    last_updated TIMESTAMP DEFAULT now()
);

INSERT INTO users(first_name, last_name, email, birthday, height, sex)
VALUES('John', 'Doe', 'john@doe.com', '2000-02-02', 182, 'Male');

SELECT * FROM users;