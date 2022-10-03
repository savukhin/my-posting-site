CREATE TABLE IF NOT EXISTS files {
    id SERIAL NOT NULL,
    filepath VARCHAR(200) NOT NULL,
    title VARCHAR(200) NOT NULL,
    file_type fileType NOT NULL,
    owner_type ownerType NOT NULL,
    owner_id INTEGER,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP,

    CONSTRAINT File_PK PRIMARY KEY(id),
}