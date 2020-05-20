CREATE TABLE IF NOT EXISTS regatta.races (
    uuid UUID PRIMARY KEY,
    name TEXT,
    start TIMESTAMP,
    distance INT
);