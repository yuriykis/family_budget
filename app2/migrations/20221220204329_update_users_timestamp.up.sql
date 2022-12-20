ALTER TABLE
    users
ADD
    COLUMN created_at timestamp NOT NULL DEFAULT now(),
ADD
    COLUMN updated_at timestamp NOT NULL DEFAULT now();