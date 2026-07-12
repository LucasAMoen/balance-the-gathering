# Infrastructure

## Postgres Setup

- Copy the example.env file within this directory into .env and fill with your desired credentials

- Run the following command:
``` bash
#apps/api/infrastructure/adapters/postgresql
docker compose up
```

- Then you can visit localhost:8080 to view the database
    - Log in with the credentials in your .env file and use db as the server

- Ensure you have sqlc and goose installed on your device

``` bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
```

-  Then run the following commands
``` bash
#apps/api/infrastructure
goose up
goose -dir ./infrastructure/adapters/postgresql/migrations/ -table _db_version up
# If you want to seed the database
goose -dir ./infrastructure/adapters/postgresql/seeds/ -table _db_seeds up
```
- This will migrate your database and seed the data

## Changes To the Database (goose)

To make changes to the database please add a migration to the [migrations](./adapters/postgresql/migrations/) directory, do so by running this command:

``` bash
#apps/api
goose -dir ./infrastructure/adapters/postgresql/migrations/ create -s test sql
```

Then if necessary create a seed for that migration

``` bash
#apps/api
goose -dir ./infrastructure/adapters/postgresql/seeds create -s test sql
```

The migration/seed names should be descriptive of what the purpose is, please see other migrations/seeds for examples

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
