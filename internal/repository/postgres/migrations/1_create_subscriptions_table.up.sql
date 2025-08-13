CREATE EXTENSION IF NOT EXISTS btree_gist;

CREATE TABLE IF NOT EXISTS subscriptions (
    id SERIAL PRIMARY KEY,
    public_id UUID NOT NULL UNIQUE DEFAULT gen_random_uuid(),
    service_name TEXT NOT NULL,
    price INTEGER NOT NULL,
    user_id UUID NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE
);

ALTER TABLE subscriptions
ADD CONSTRAINT valid_price_value
CHECK (price >= 0);

ALTER TABLE subscriptions
ADD CONSTRAINT end_date_after_start_date
CHECK (end_date IS NULL OR end_date > start_date);

ALTER TABLE subscriptions
ADD CONSTRAINT no_overlapping_subscriptions
EXCLUDE USING GIST (
    user_id WITH =,
    service_name WITH =,
    daterange(start_date, end_date, '[)') WITH &&
);