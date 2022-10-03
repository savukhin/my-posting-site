DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'ownerType') THEN
        CREATE TYPE ownerType AS ENUM 
        (
            'user',
            'post'
        );
    END IF;
END$$;