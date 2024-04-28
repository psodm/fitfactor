CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "citext";

DROP TABLE IF EXISTS user_profile;
DROP TABLE IF EXISTS user_account CASCADE;
DROP TYPE IF EXISTS sex;
DROP TYPE IF EXISTS unit;

CREATE TYPE sex AS ENUM ('Male', 'Female');
CREATE TYPE unit AS ENUM('Imperial', 'Metric');

CREATE TABLE IF NOT EXISTS user_account (
    id UUID DEFAULT uuid_generate_v4(),
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
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

/* password */
INSERT INTO user_account(username, email, password)
VALUES('John Doe', 'john@doe.com', '$2a$12$SwOvMIS0TF/KM5OMCKEGEuepXCCUnY/USFi3mAts1UPLKpJRIgYUK');

/* password123 */
INSERT INTO user_account(username, email, password)
VALUES('Jane Doe', 'jane@doe.com', '$2a$12$oCNJ8mF4pOCyk6iAHba7AOGZ.4g90sipVyFGKcNRGSlaJ8Q2PjiwK');

INSERT INTO user_profile(user_id, date_of_birth, sex)
VALUES(
    (SELECT id FROM user_account WHERE email = 'john@doe.com'),
    '2000-02-02',
    'Male'
);

INSERT INTO user_profile(user_id, date_of_birth, sex)
VALUES(
    (SELECT id FROM user_account WHERE email = 'jane@doe.com'),
    '2001-01-01',
    'Female'
);

SELECT * FROM user_account;

SELECT * FROM user_profile;