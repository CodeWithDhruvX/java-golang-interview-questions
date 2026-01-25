# Composite Pattern

## 🟢 What is it?
The **Composite Pattern** lets you compose objects into tree structures and then work with these structures as if they were individual objects.

Think of it like a **File System**:
*   You have **Files** (Leafs) and **Folders** (Composites).
*   A Folder can contain Files OR other Folders.
*   If you ask "What is the size of this Folder?", the Folder asks all items inside it for their size and sums them up.
*   You treat a File and a Folder exactly the same way (using `getSize()`).

---

## 🎯 Strategy to Implement

1.  **Component Interface**: Declare the interface that is common to both simple and complex objects (e.g., `FileSystemItem` with `getSize()`).
2.  **Leaf Class**: basic element that doesn't have sub-elements (e.g., `File`). Implements `getSize()` by returning its own size.
3.  **Composite Class**: Element that has sub-elements (e.g., `Folder`). Stores a list of child Components. Implements `getSize()` by iterating over children and summing up results.

---

## 💻 Code Example

```java
import java.util.ArrayList;
import java.util.List;

// 1. Component Interface
interface FileSystemItem {
    void print(String temp);
}

// 2. Leaf Class
class File implements FileSystemItem {
    private String name;

    public File(String name) {
        this.name = name;
    }

    @Override
    public void print(String structure) {
        System.out.println(structure + "File: " + name);
    }
}

// 3. Composite Class
class Folder implements FileSystemItem {
    private String name;
    private List<FileSystemItem> children = new ArrayList<>();

    public Folder(String name) {
        this.name = name;
    }

    public void add(FileSystemItem item) {
        children.add(item);
    }

    @Override
    public void print(String structure) {
        System.out.println(structure + "Folder: " + name);
        for (FileSystemItem child : children) {
            // Recursive call
            child.print(structure + "   ");
        }
    }
}
```

### Usage:

```java
public class Main {
    public static void main(String[] args) {
        File f1 = new File("Song.mp3");
        File f2 = new File("Picture.png");
        File f3 = new File("Doc.pdf");

        Folder music = new Folder("Music");
        music.add(f1);

        Folder docs = new Folder("Documents");
        docs.add(f2);
        docs.add(f3);
        docs.add(music); // Adding folder inside folder

        // Print the whole tree structure
        docs.print("");
    }
}
```

---

## ✅ When to use?

*   **Tree Structures**: When you need to implement a tree-like object structure (e.g., Organization Chart, Menus, File Systems, GUI Elements).
*   **Uniformity**: When you want the client code to treat both simple and complex elements uniformly. The client doesn't need to check `if (object is folder) ... else ...`, it just calls `object.operation()`.
