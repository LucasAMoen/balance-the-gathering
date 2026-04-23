# Balance The Gathering

## Joining the Game
- Decks joining the game must begin less than 15 dollars
- Basic lands do not count toward the cost
- Anything other than cards (counters, dice, deck box, sleeves) do not count toward the cost
- Packs can be purchased and all cards you get from the pack can be used
- Anything you can get with 15 dollars

## Points
- Everyone starts with 0 points
- Points are worth about 1 cent each

### Bounties
- Bounties can be put on other players
- A bounty is paid for in points and a player who collects a bounty gets the points
- The player to deliver the final blow collects the bounty

### Earning points
- Winning a 4 pod earns 500 points
- Winning a 3 pod earns 250 points
- Winning a 1v1 earns 100 points

### Purchasing Cards
- Players can use/purchase/print/create cards with their points
- 1 point is equivalent to 1 cent
- If players want to add a new card to their deck they must pay the cards cost in points via the application
- Price is determined by an official api

### Killing Cards
- Players can "kill" cards other people own if both of the following are true
  - The card is in the graveyard/exile
  - AND
  - Twice the cost of the card in points must be paid
- Killing is a state based action and will be tracked
- The card can be used as normal for the rest of the game
- Once the kill condition is met, the card is killed and must be purchased again by the player with points
- Cards are only killed after a game is finished
- Basic lands cannot be killed
- Non-basic lands can be killed
- Commander can be killed
- Any amount of players can put in points to kill a card

### Graphical Interfaces
- Everyone logs in on their phone and has their life total on their own phone
- Everyone must be logged in and in the lobby before the game can start
- Once a game starts everyone begins at 40 health
- Commander damage can be taken in the app
- Bounties can be placed any time within or outside a game

### Application
- Uses the scryfall api to find cards
- Users import their decks usingstandard mtg export/import

# React + TypeScript + Vite

## Expanding the ESLint configuration

If you are developing a production application, we recommend updating the configuration to enable type-aware lint rules:

```js
export default defineConfig([
  globalIgnores(['dist']),
  {
    files: ['**/*.{ts,tsx}'],
    extends: [
      // Other configs...

      // Remove tseslint.configs.recommended and replace with this
      tseslint.configs.recommendedTypeChecked,
      // Alternatively, use this for stricter rules
      tseslint.configs.strictTypeChecked,
      // Optionally, add this for stylistic rules
      tseslint.configs.stylisticTypeChecked,

      // Other configs...
    ],
    languageOptions: {
      parserOptions: {
        project: ['./tsconfig.node.json', './tsconfig.app.json'],
        tsconfigRootDir: import.meta.dirname,
      },
      // other options...
    },
  },
])
```

```js
// eslint.config.js
import reactX from 'eslint-plugin-react-x'
import reactDom from 'eslint-plugin-react-dom'

export default defineConfig([
  globalIgnores(['dist']),
  {
    files: ['**/*.{ts,tsx}'],
    extends: [
      // Other configs...
      // Enable lint rules for React
      reactX.configs['recommended-typescript'],
      // Enable lint rules for React DOM
      reactDom.configs.recommended,
    ],
    languageOptions: {
      parserOptions: {
        project: ['./tsconfig.node.json', './tsconfig.app.json'],
        tsconfigRootDir: import.meta.dirname,
      },
      // other options...
    },
  },
])
```
