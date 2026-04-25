# Supabase Application

Supabase is the auth/database

Default readme is [here](/apps/supabase/DEFAULTREADME.md)

## Postgres Setup

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
