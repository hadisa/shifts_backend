-- Create uuid v4 extenstion
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
    id uuid,
    is_superuser boolean,
    email character varying(254) NOT NULL,
    phone character varying(16),
    whatsapp character varying(16),
    is_staff boolean NOT NULL,
    is_active boolean NOT NULL,
    date_joined timestamp with time zone NOT NULL,
    last_login timestamp with time zone,
    note text,
    first_name character varying(256) NOT NULL,
    last_name character varying(256) NOT NULL,
    avatar character varying(100),
    private_metadata jsonb,
    metadata jsonb,
    language_code character varying(35) NOT NULL,
    search_document text,
    updated_at timestamp with time zone NOT NULL,
    PRIMARY KEY (id)
);

