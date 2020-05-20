CREATE TABLE IF NOT EXISTS regatta.regattas (
    id UUID PRIMARY KEY,
    name TEXT,
    start TIMESTAMP,
    stop TIMESTAMP,
    races SET<FROZEN<flyweight>>
);