CREATE TABLE IF NOT EXISTS regatta.teams (
    id UUID PRIMARY KEY ,
    name TEXT,
    users SET<FROZEN<flyweight>>
);