# Low-Level Design (LLD) - Chess Game

## Problem Statement
Design a classic 2-player board game of Chess. The system should manage turns, pieces, and enforce game rules (valid moves, check, checkmate, castling).

## Requirements
*   **Board:** An 8x8 grid.
*   **Pieces:** 16 pieces per player (King, Queen, Rooks, Bishops, Knights, Pawns).
*   **Move Validation:** Each piece has specific move rules.
*   **Game State:** Must detect Check, Checkmate, Stalemate, and Draw conditions.
*   **Special Moves:** En passant, castling, and pawn promotion.
*   **History:** Ability to track moves and undo (optional but good).

## Core Entities / Classes

1.  **Game:** Central controller. Has 2 `Player`s, a `Board`, `currentTurn`, and `GameStatus` (ACTIVE, BLACK_WON, WHITE_WON, STALEMATE).
2.  **Board:** 8x8 array of `Box` (or `Square`).
3.  **Box / Square:** Contains `color` (Black/White), `x`, `y` coordinates, and optionally a `Piece`.
4.  **Player:** Account details, Color (White/Black).
5.  **Piece (Abstract):** `killed`, `white` (boolean). Abstract method `canMove(Board board, Box start, Box end)`.
6.  **Concrete Pieces (King, Queen, Rook, etc.):** Implement `canMove()` overriding custom logic.
7.  **Move:** Represents a single move: `Player`, `startBox`, `endBox`, `pieceMoved`, `pieceKilled`.

## Key Design Patterns Applicable
*   **Command Pattern:** Essential for the `Move` object. It encapsulation actions and makes it incredibly easy to implement an undo/redo feature, or replay an entire game.
*   **Factory Pattern:** To set up the board and initialize the 32 pieces in their correct `Box`.
*   **Strategy Pattern:** Move validation (often implicit via Polymorphism in the Piece abstract class).

## Code Snippet (Polymorphism and Move Validation)

```java
public abstract class Piece {
    private boolean isWhite;
    private boolean isKilled;

    public abstract boolean canMove(Board board, Box start, Box end);
    // getters and setters
}

public class Knight extends Piece {
    @Override
    public boolean canMove(Board board, Box start, Box end) {
        // Can't move to a box with own piece
        if (end.getPiece() != null && end.getPiece().isWhite() == this.isWhite()) {
            return false;
        }

        int x = Math.abs(start.getX() - end.getX());
        int y = Math.abs(start.getY() - end.getY());
        return (x * y == 2); // 'L' shape: 1*2 or 2*1
    }
}
```

## Follow-up Questions for Candidate
1.  How do you implement the "Checkmate" detection efficiently without simulating all possible moves heavily?
2.  How would you design an AI bot to play? (Minimax algorithm + Alpha-Beta pruning)
3.  Explain how you handle "Pawn Promotion". Does it require replacing the Piece instance on the Board?
