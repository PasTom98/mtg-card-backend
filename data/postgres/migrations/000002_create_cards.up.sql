CREATE TABLE IF NOT EXISTS "cards" (
    id bigserial PRIMARY KEY,
    scryfall_id uuid NOT NULL,
    name varchar(255) NOT NULL,
    oracle_text text NOT NULL,
    mana_cost varchar(255) NOT NULL,
    types text[] NOT NULL,
    set varchar(255) NOT NULL
    );