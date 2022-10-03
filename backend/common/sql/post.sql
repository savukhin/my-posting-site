CREATE TABLE IF NOT EXISTS posts {
    id SERIAL NOT NULL,
    author_id INTEGER,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP,

    CONSTRAINT Post_PK PRIMARY KEY(id),
}