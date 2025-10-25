CREATE TABLE metadata (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT,
    size INT NOT NULL,
    type TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);
