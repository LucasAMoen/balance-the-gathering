# Api Application

## Bulk Download

- Lazy Bulk download from scryfall if both are true:
    - User bulk uploads their deck/searches/purchases a card
    - and
    - It has been > 1 day since last bulk upload
- OR
- Whenever someone starts the server
- Bulk pull json from scryfall of all card data
- Dump all card data into the Postgres "Cards" table
- Keep track of last bulk upload via local cache
