CREATE TABLE boards (
    slug TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT
);

CREATE TABLE threads (
    id SERIAL PRIMARY KEY,
    board_slug TEXT NOT NULL REFERENCES boards(slug),
    title TEXT,
    author TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    sticky BOOLEAN NOT NULL DEFAULT FALSE,
    locked BOOLEAN NOT NULL DEFAULT FALSE,
    bump_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    thread_id INTEGER NOT NULL REFERENCES threads(id) ON DELETE CASCADE,
    board_slug TEXT NOT NULL REFERENCES boards(slug),
    author TEXT,
    text TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    image_url TEXT,
    parent_id INTEGER REFERENCES posts(id),
    tripcode TEXT
);