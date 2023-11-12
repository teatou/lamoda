CREATE TABLE IF NOT EXISTS item (
    id SERIAL UNIQUE,
    title VARCHAR,
    code VARCHAR UNIQUE,
    reserved BOOLEAN
)