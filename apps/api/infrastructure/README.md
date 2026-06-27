# Supabase Application

Postgres is the database

## Postgres Setup

- Copy the example.env file into .env and fill with your desired credentials

- Run the following command:
``` bash

#apps/api/infrastructure/adapters/postgresql
docker compose up

```

- Then you can visit localhost:8080 to view the database
    - Log in with the credentials in your .env file and use db as the server


User:
- id: uuid
- username: string
- points: int
- bounty: fk to bounty
- decks: array fk to deck id
- activeGame: fk to game
- createdAt: date

Deck:
- id: uuid
- commander: fk to card
- cards: array fk to card
- createdAt: date

OwnedCards:
- id: uuid
- user: fk to user
- card: fk to cards
- isKilled: bool
- priceAtPurchase: decimal
- obtainedAt: date

Cards:
- id: scryfall id
- name: string
- imageUrl: url string
- price: decimal
- createdAt: date

Game
- id: uuid
- joinCode: 4 characters (easy to remember)
- players: array fk to user
- pointAverage: int
- commanderDamage: fk to CommanderDamage
- winner: fk to user
- gameStatus: enum (inProgress, finished, tied)
- gameStart: date
- gameFinish: date
- createdAt: date

CommanderDamage:
- id: uuid
- dealingPlayer: fk to user
- receivingPlayer: fk to user
- damage: int
- dealtAt: date

Bounty:
- id: uuid
- givingPlayer?: fk to user
- targetPlayer: fk to user
- amount: int
- collectedAt?: date
- createdAt: date

BountyCollectionRequest
- id: uuid
- requestingPlayer: fk to user
- targetPlayer: fk to user
- confirmed: bool
- createdAt: date

CardKillRequest:
- id: uuid
- requestingPlayer: fk to user
- targetPlayer: fk to user
- paidFrom: fk array to user (include amount paid)
- confirmed: bool
- createdAt: date
