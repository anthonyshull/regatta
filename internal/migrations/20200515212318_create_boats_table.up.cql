CREATE TABLE IF NOT EXISTS regatta.boats (
    id UUID PRIMARY KEY,
    name TEXT,
    team flyweight,
    users SET<FROZEN<flyweight>>
);