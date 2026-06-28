-- +goose Up
CREATE TABLE IF NOT EXISTS cards (
    id UUID PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    layout TEXT,
    imageUrl TEXT,
    price DECIMAL NOT NULL CHECK (price >= 0),
    createdAt TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
    updatedAt TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS cards;
