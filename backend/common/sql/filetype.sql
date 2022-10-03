DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'fileType') THEN
        CREATE TYPE content_types AS ENUM 
        (
            'photo',
            'text', 
        );
    END IF;
END$$;