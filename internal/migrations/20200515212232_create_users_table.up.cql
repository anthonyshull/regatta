CREATE TABLE IF NOT EXISTS regatta.users (
    id UUID PRIMARY KEY,
    name TEXT,
    email TEXT,
    password TEXT,
    salt TEXT,
    teams SET<FROZEN<flyweight>>
);