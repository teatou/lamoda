CREATE TABLE IF NOT EXISTS warehouse (
    id SERIAL UNIQUE,
    title VARCHAR,
    available BOOLEAN
)