CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS user_characteristic;
DROP TABLE IF EXISTS user_account CASCADE;
DROP TYPE IF EXISTS sex;
DROP TYPE IF EXISTS unit;

CREATE TYPE sex AS ENUM ('Male', 'Female');
CREATE TYPE unit AS ENUM('Imperial', 'Metric');

CREATE TABLE IF NOT EXISTS user_account (
    id UUID DEFAULT uuid_generate_v4(),
    username TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL,
    date_created TIMESTAMP DEFAULT now(),
    last_updated TIMESTAMP DEFAULT now(),
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS user_profile (
    user_id UUID,
    date_of_birth DATE,
    sex sex,
    profile_image_url TEXT,
    preferred_unit unit DEFAULT 'Metric',
    PRIMARY KEY(user_id),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES user_account(id)
);

INSERT INTO user_account(username, email, hashed_password)
VALUES('John Doe', 'john@doe.com', '$2a$12$SwOvMIS0TF/KM5OMCKEGEuepXCCUnY/USFi3mAts1UPLKpJRIgYUK');

INSERT INTO user_profile(user_id, date_of_birth, sex)
VALUES(
    (SELECT id FROM user_account WHERE email = 'john@doe.com'),
    '2000-02-02',
    'Male'
);

SELECT * FROM user_account;

SELECT * FROM user_profile;