# Composite Pattern

## 🟢 What is it?
The **Composite Pattern** lets you compose objects into tree structures and then work with these structures as if they were individual objects.

Think of it like a **File System**:
*   You have **Files** (Leafs) and **Folders** (Composites).
*   A Folder can contain Files OR other Folders.
*   You treat a File and a Folder exactly the same way (using `GetSize()`).

---

## 🎯 Strategy to Implement

1.  **Component Interface**: Declare the interface that is common to both simple and complex objects.
2.  **Leaf Struct**: Basic element that doesn't have sub-elements.
3.  **Composite Struct**: Element that has sub-elements. Stores a list of child Components. Implements the methods by iterating over children.

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. Component Interface
type FileSystemItem interface {
    Print(indent string)
}

// 2. Leaf Struct
type File struct {
    Name string
}

func (f *File) Print(indent string) {
    fmt.Println(indent + "File: " + f.Name)
}

// 3. Composite Struct
type Folder struct {
    Name     string
    Children []FileSystemItem
}

func (f *Folder) Add(item FileSystemItem) {
    f.Children = append(f.Children, item)
}

func (f *Folder) Print(indent string) {
    fmt.Println(indent + "Folder: " + f.Name)
    for _, child := range f.Children {
        child.Print(indent + "   ")
    }
}

func main() {
    f1 := &File{Name: "Song.mp3"}
    f2 := &File{Name: "Picture.png"}
    f3 := &File{Name: "Doc.pdf"}

    music := &Folder{Name: "Music"}
    music.Add(f1)

    docs := &Folder{Name: "Documents"}
    docs.Add(f2)
    docs.Add(f3)
    docs.Add(music) // Adding folder inside folder

    // Print the whole tree structure
    docs.Print("")
}
```

---

## ✅ When to use?

*   **Tree Structures**: When you need to implement a tree-like object structure (e.g., Organization Chart, Menus, File Systems, GUI Elements).
*   **Uniformity**: When you want the client code to treat both simple and complex elements uniformly.
