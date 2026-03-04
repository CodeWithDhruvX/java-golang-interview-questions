# Low-Level Design (LLD) - Library Management System

## Problem Statement
Design a Library Management System (LMS) that allows users to search for books, checkout books, return books, and pay early or late fees.

## Requirements
*   **Users:** Two types of users: Librarian and Member.
*   **Books:** A book can have multiple copies (BookItems). A book has properties like Title, Subject, Author.
*   **Operations:** 
    *   Search by title, author, subject, or publication date.
    *   Checkout a book (limit max books per user).
    *   Return a book.
    *   Reserve a book if currently not available.
*   **Fines:** The system must calculate and charge a fine for overdue books.
*   **Notifications:** Send alerts for reserved books becoming available or for overdue fees.

## Core Entities / Classes

1.  **LMS (Singleton):** Central system object.
2.  **Book:** Represents the catalog details (ISBN, Title, Subject, Authors).
3.  **BookItem:** A physical copy of a `Book`. Has a `barcode`, `isReferenceOnly`, `price`, and `Status` (AVAILABLE, RESERVED, LOANED, LOST).
4.  **Account (Abstract):** `id`, `password`, `status` (ACTIVE, CLOSED, CANCELED).
    *   `Librarian`: Can add/remove books, block members.
    *   `Member`: Can checkout, return, reserve, pay fines.
5.  **BookReservation:** Details regarding reserving a BookItem.
6.  **BookLending:** Tracks the checkout. Has `creationDate`, `dueDate`, `returnDate`, and `associatedMember`.
7.  **Fine:** Calculates fines based on `BookLending`.
8.  **Search (Interface):** Defines search capabilities (ByTitle, ByAuthor).

## Key Design Patterns Applicable
*   **Factory Pattern:** To create different types of Accounts (Librarian, Member).
*   **Observer Pattern:** When a book becomes AVAILABLE, notify the members who have RESERVED it.
*   **Strategy Pattern:** For search algorithms. A Member can choose a strategy (SearchByTitle, SearchByAuthor).

## Code Snippet (Search Strategy)

```java
public interface Search {
    public List<Book> searchByTitle(String title);
    public List<Book> searchByAuthor(String author);
}

public class Catalog implements Search {
    private HashMap<String, List<Book>> bookTitles;
    private HashMap<String, List<Book>> bookAuthors;

    public List<Book> searchByTitle(String title) {
        return bookTitles.getOrDefault(title, new ArrayList<>());
    }

    public List<Book> searchByAuthor(String author) {
        return bookAuthors.getOrDefault(author, new ArrayList<>());
    }
}
```

## Follow-up Questions for Candidate
1.  How will the system handle concurrency if two members try to checkout the very last copy of a book at the exact same moment?
2.  How would you design the database schema to normalize Authors and Books given many-to-many relationships?
3.  How can you implement an automated daily job to identify overdue books and calculate initial fine entries?
