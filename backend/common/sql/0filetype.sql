DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'fileType') THEN
        CREATE TYPE fileType AS ENUM 
        (
            'photo',
            'text'
        );
    END IF;
END$$;