CREATE TABLE IF NOT EXISTS regatta.results (
    id UUID PRIMARY KEY,
    race flyweight,
    boat flyweight,
    place INT,
    time DURATION
);