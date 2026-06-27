# Balance The Gathering

## Contents

- Other Documents:
  - [Developer Readme](/DEVELOP.md)

- This Document:
  - [Game Rules](#game-rules)
  - [AI Use Statement](#ai-use-statement)

## Game Rules

### Joining the Game
- Decks joining the game must begin less than 15 dollars
- Basic lands do not count toward the cost
- Anything other than cards (counters, dice, deck box, sleeves) do not count toward the cost
- Packs can be purchased and all cards you get from the pack can be used
- Anything you can get with 15 dollars

### Points
- Everyone starts with 0 points
- Points are worth about 1 cent each

#### Bounties
- Bounties can be put on other players
- A bounty is paid for in points and a player who collects a bounty gets the points
- The player to deliver the final blow collects the bounty
- Players start to rack up bounties automatically if they have won multiple games in a row
  - 2 games in a row - 100 bounty points
  - 3 games in a row - 200 bounty points
  - 4 games in a row - 300 bounty points
  - 5 games in a row - 400 bounty points
  - 6 games in a row - 500 bounty points
  - Maxes out at 500 bounty points

#### Earning points
- Winning a 4 pod earns 500 points
- Winning a 3 pod earns 250 points
- Winning a 1v1 earns 100 points

#### ELO System:
- When a user joins the game, the average of everyone's total point value will be calculated.
- When someone wins the game, if they had less than that average total, they will receive points equal to 50% the difference

#### Purchasing Cards
- Players can use/purchase/print/create cards with their points
- 1 point is equivalent to 1 cent
- If players want to add a new card to their deck they must pay the cards cost in points via the application
- Price is determined by an official api
- When players purchase new cards, they still own their old ones which can be used to build other decks

#### Killing Cards
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

#### Graphical Interfaces
- Everyone logs in on their phone and has their life total on their own phone
- Everyone must be logged in and in the lobby before the game can start
- Once a game starts everyone begins at 40 health
- Commander damage can be taken in the app
- Bounties can be placed any time within or outside a game

## AI Use Statement

The content in this codebase contains **absolutely zero** AI generated code. This game is meant to be an artisan code project to hone in my programming skills and passion for software engineering. The **only** times AI have been used in this project are for outstanding bugs related to setup and/or issues not solvable by parsing through documentation. It is a slower process, but it allows me to fully learn the content without a crutch and gives me the depth of experience applicable to any other project.