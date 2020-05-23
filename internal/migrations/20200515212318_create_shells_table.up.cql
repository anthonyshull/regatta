CREATE TABLE IF NOT EXISTS regatta.shells (
    id UUID PRIMARY KEY,
    name TEXT,
    team flyweight,
    users SET<FROZEN<flyweight>>
);