# Low-Level Design (LLD) - Snake and Ladder

## Problem Statement
Design a classic Board Game: Snake and Ladder. Multiple players can play the game on a standard 10x10 board.

## Requirements
*   **Board:** A board with sizes like 10x10 (1 to 100).
*   **Entities:** Players, Snakes, Ladders, and Dice.
*   **Rules:**
    *   The game is played sequentially.
    *   Players roll a dice (1-6) to move.
    *   If a player lands on the start of a Ladder, they go up.
    *   If a player lands on the head of a Snake, they go down.
    *   A player wins if they land exactly on 100.
*   **Optional Enhancements:** Multiple dice, special game rules (like getting another turn on a 6).

## Core Entities / Classes

1.  **Game:** Central controller. Contains a Queue of `Player`s, the `Board`, and the `Dice`. Determines the game loop.
2.  **Board:** Holds sizes (e.g., 100 cells) and a list of `Snake`s and `Ladder`s.
3.  **Player:** Has an `id`, `name`, and `currentPosition`.
4.  **Dice:** Has `numberOfDice` and a `roll()` method utilizing secure random generation.
5.  **Jumper (Abstract):** Represents elements that instantly move a player.
    *   `Snake` (inherits Jumper): `start` > `end`.
    *   `Ladder` (inherits Jumper): `start` < `end`.

## Key Design Patterns Applicable
*   **Queue Data Structure:** A `java.util.Queue` is perfect for managing player turns sequentially. Dequeue the player, roll the dice, update position, and enqueue them back if they didn't win.
*   **Factory Pattern:** Generating the board setup randomly with a specific number of snakes and ladders.

## Code Snippet (Game Loop & Queue logic)

```java
public class Game {
    private Queue<Player> players;
    private Board board;
    private Dice dice;

    public void play() {
        while (players.size() > 1) { // Stop when 1 player is remaining
            Player currentPlayer = players.poll();
            int diceVal = dice.roll();
            int newPos = currentPlayer.getCurrentPosition() + diceVal;

            if (newPos > board.getSize()) {
                players.offer(currentPlayer); // Turn wasted, wait for exact roll
                continue;
            }

            // Check for snakes/ladders
            newPos = board.getFinalPositionAfterJump(newPos);
            currentPlayer.setCurrentPosition(newPos);

            if (currentPlayer.getCurrentPosition() == board.getSize()) {
                System.out.println(currentPlayer.getName() + " won the game!");
            } else {
                players.offer(currentPlayer); // Put back in queue
            }
        }
    }
}
```

## Follow-up Questions for Candidate
1.  How do you ensure you don't create an infinite loop where the tail of a snake drops you on the start of a ladder that takes you right back up to the snake? 
2.  How would you modify this design to support a completely generic board game engine (like plugging in Monopoly rules vs Snake and Ladder rules)?
