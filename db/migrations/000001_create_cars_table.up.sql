CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE cars (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    model TEXT NOT NULL,
    color TEXT NOT NULL,
    year TEXT NOT NULL,
    price TEXT NOT NULL,
    image TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL
);
