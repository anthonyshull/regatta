CREATE TABLE IF NOT EXISTS regatta.races (
    id UUID PRIMARY KEY,
    name TEXT,
    start TIMESTAMP,
    distance INT,
    boats SET<FROZEN<flyweight>>
);