CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

--- Creating Tables
CREATE TABLE links
(
    id           uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    url          TEXT NOT NULL UNIQUE,
    website      TEXT NOT NULL,
    checked_date DATE NOT NULL
);