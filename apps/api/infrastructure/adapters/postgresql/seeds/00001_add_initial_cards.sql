-- +goose Up
INSERT INTO
cards(
    id,
    name,
    url,
    layout,
    imageUrl,
    price,
    createdAt,
    updatedAt

)
VALUES
(
    '0000419b-0bba-4488-8f7a-6194544ce91e',
    'Forest',
    'https://api.scryfall.com/cards/0000419b-0bba-4488-8f7a-6194544ce91e',
    'normal',
    'https://cards.scryfall.io/png/front/0/0/0000419b-0bba-4488-8f7a-6194544ce91e.png?1721427487',
    0.32,
    now(),
    now()
),
(
    'e4a2d2c6-8eaa-4760-b620-921b807baa2e',
    'Feather, the Redeemed',
    'https://api.scryfall.com/cards/e4a2d2c6-8eaa-4760-b620-921b807baa2e',
    'normal',
    'https://cards.scryfall.io/png/front/e/4/e4a2d2c6-8eaa-4760-b620-921b807baa2e.png?1557577142',
    0.71,
    now(),
    now()
);

-- +goose Down
DELETE FROM cards WHERE
id='0000419b-0bba-4488-8f7a-6194544ce91e'
or
id='e4a2d2c6-8eaa-4760-b620-921b807baa2e';