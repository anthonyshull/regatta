CREATE TABLE IF NOT EXISTS regatta.results (
    id UUID PRIMARY KEY,
    race flyweight,
    shell flyweight,
    place INT,
    time DURATION
);