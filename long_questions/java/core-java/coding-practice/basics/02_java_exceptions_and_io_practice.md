# ⚠️ Java Exceptions & IO Practice
Contains runnable code examples for Questions 21-40.

## Question 21: Exception Hierarchy in Java.

### Answer
Throwable -> Error / Exception (Checked / Unchecked).

### Runnable Code
```java
package basics;

import java.io.IOException;

public class ExceptionHierarchy {
    public static void main(String[] args) {
        // 1. Error (Unrecoverable)
        try {
            recursive();
        } catch (StackOverflowError e) {
            System.out.println("Caught Error: " + e.getClass().getSimpleName());
        }
        
        // 2. RuntimeException (Unchecked)
        try {
            String s = null;
            s.length();
        } catch (NullPointerException e) {
            System.out.println("Caught Unchecked: " + e);
        }
        
        // 3. Checked Exception (Must Handle)
        try {
            throwChecked();
        } catch (IOException e) {
            System.out.println("Caught Checked: " + e.getMessage());
        }
    }
    
    static void recursive() { recursive(); }
    
    static void throwChecked() throws IOException {
        throw new IOException("File not found simulation");
    }
}
```

---

## Question 22: Validating `try-catch`, `finally` combinations.

### Answer
`try-catch`, `try-finally`, `try-catch-finally`. `try` alone is invalid.

### Runnable Code
```java
package basics;

public class TryCatchValidations {
    public static void main(String[] args) {
        // Case 1: Try-Catch
        try {
            int x = 10/0;
        } catch (ArithmeticException e) {
            System.out.println("1. Caught / by zero");
        }
        
        // Case 2: Try-Finally (Exceptions propagate up if not caught)
        try {
            System.out.println("2. Try Block");
        } finally {
            System.out.println("2. Finally Block");
        }
        
        // Case 3: Try-Catch-Finally
        try {
            System.out.println("3. Working...");
        } catch (Exception e) {
            // Not executed
        } finally {
            System.out.println("3. Cleanup");
        }
    }
}
```

---

## Question 23: `throw` vs `throws`.

### Answer
`throw`: Action inside method. `throws`: Declaration in signature.

### Runnable Code
```java
package basics;

public class ThrowVsThrows {
    // 'throws' declares that this method MIGHT fail
    static void checkAge(int age) throws Exception {
        if (age < 18) {
            // 'throw' explicitly creates the exception
            throw new Exception("Access Denied");
        }
        System.out.println("Access Granted");
    }

    public static void main(String[] args) {
        try {
            checkAge(15);
        } catch (Exception e) {
            System.out.println("Caught: " + e.getMessage());
        }
    }
}
```

---

## Question 24: `final`, `finally`, `finalize`.

### Answer
`final` (const), `finally` (cleanup block), `finalize` (GC method - deprecated).

### Runnable Code
```java
package basics;

public class FinalKeywords {
    // 1. Final Variable
    final int MAX = 100;

    // 1. Final Method
    final void noOverride() {}

    public static void main(String[] args) {
        FinalKeywords obj = new FinalKeywords();
        // obj.MAX = 200; // Error
        
        // 2. Finally
        try {
            System.out.println("Try");
        } finally {
            System.out.println("Finally always runs");
        }
        
        // 3. Finalize (Simulated trigger)
        obj = null;
        System.gc();
    }
    
    @Override
    protected void finalize() throws Throwable {
        System.out.println("Finalize called by GC (Deprecated)");
    }
}
```

---

## Question 25: What is Try-with-Resources? (Java 7+)

### Answer
Auto-closes `AutoCloseable` resources.

### Runnable Code
```java
package basics;

// Custom Resource
class MyResource implements AutoCloseable {
    @Override
    public void close() {
        System.out.println("Resource closed automatically!");
    }
    
    public void read() { System.out.println("Reading resource..."); }
}

public class TryWithResources {
    public static void main(String[] args) {
        // Resources declared in () are closed at end of block
        try (MyResource res = new MyResource()) {
            res.read();
        } // .close() called here implicitly
    }
}
```

---

## Question 26: Checked vs Unchecked Exception? (When to use?)

### Answer
Checked (recoverable, external), Unchecked (programming error).

### Runnable Code
```java
package basics;

import java.io.FileNotFoundException;

public class CheckedVsUnchecked {
    // Checked: Forces caller to handle (e.g., config missing)
    static void readFile() throws FileNotFoundException {
        // throw new FileNotFoundException("config.txt missing");
    }

    // Unchecked: Caller ignores (e.g., bad logic)
    static void divide(int a, int b) {
        if (b == 0) throw new IllegalArgumentException("Divisor cannot be 0");
        System.out.println(a / b);
    }

    public static void main(String[] args) {
        try {
            readFile();
        } catch (FileNotFoundException e) {
            System.out.println("Handled Checked Ex");
        }
        
        divide(10, 2); // No try-catch forced
    }
}
```

---

## Question 27: Custom Exception creation.

### Answer
Extend `Exception` or `RuntimeException`.

### Runnable Code
```java
package basics;

class InsufficientFundsException extends RuntimeException {
    public InsufficientFundsException(String msg) {
        super(msg);
    }
}

class BankAccount {
    int balance = 100;
    
    void withdraw(int amount) {
        if (amount > balance) {
            throw new InsufficientFundsException("Need " + amount + " but have " + balance);
        }
        balance -= amount;
        System.out.println("Withdraw success. Balance: " + balance);
    }
}

public class CustomExceptionDemo {
    public static void main(String[] args) {
        try {
            new BankAccount().withdraw(150);
        } catch (InsufficientFundsException e) {
            System.err.println("Transaction Failed: " + e.getMessage());
        }
    }
}
```

---

## Question 28: What happen if you throw exception in `finally` block?

### Answer
Masks the original exception.

### Runnable Code
```java
package basics;

public class ExceptionInFinally {
    public static void main(String[] args) {
        try {
            System.out.println("Start");
            try {
                throw new RuntimeException("Original Exception");
            } finally {
                // New exception hides the original
                throw new RuntimeException("Finally Exception");
            }
        } catch (Exception e) {
            System.out.println("Caught: " + e.getMessage()); // Prints "Finally Exception"
        }
    }
}
```

---

## Question 29: Exception Propagation.

### Answer
Unchecked propagates automatically. Checked must be declared.

### Runnable Code
```java
package basics;

public class PropagationDemo {
    static void m1() { m2(); }
    static void m2() { m3(); }
    static void m3() { 
        throw new RuntimeException("Boom!"); // Unchecked propagates m3 -> m2 -> m1
    }

    public static void main(String[] args) {
        try {
            m1();
        } catch (RuntimeException e) {
            System.out.println("Caught in Main: " + e.getMessage());
            e.printStackTrace(); // Shows stack trace m3 -> m2 -> m1
        }
    }
}
```

---

## Question 30: What is Serialization?

### Answer
Convert Object -> Byte Stream.

### Runnable Code
```java
package basics;

import java.io.*;

class User implements Serializable {
    String name;
    User(String n) { name = n; }
}

public class SerializationSimple {
    public static void main(String[] args) throws IOException, ClassNotFoundException {
        User user = new User("Dhruv");
        File f = new File("user.ser");
        
        // Serialize
        try (ObjectOutputStream out = new ObjectOutputStream(new FileOutputStream(f))) {
            out.writeObject(user);
        }
        
        // Deserialize
        try (ObjectInputStream in = new ObjectInputStream(new FileInputStream(f))) {
            User loaded = (User) in.readObject();
            System.out.println("Loaded: " + loaded.name);
        }
        
        f.delete(); // Cleanup
    }
}
```

---

## Question 31: `serialVersionUID` significance.

### Answer
Version control for classes.

### Runnable Code
```java
package basics;

import java.io.Serializable;

public class VersionedClass implements Serializable {
    // Explicit ID ensures class changes don't break compatibility 
    // for existing serialized files (if changes are compatible)
    private static final long serialVersionUID = 1L;
    
    String data = "Initial";
    
    // Changing this class structure without changing UID allows 
    // deserialization (fields might depend on default values).
}
```
*(Note: Code is descriptive, demonstrating the field declaration)*

---

## Question 32: `transient` keyword.

### Answer
Skip field during serialization.

### Runnable Code
```java
package basics;

import java.io.*;

class SecureUser implements Serializable {
    String username;
    transient String password; // Will not be saved
    
    SecureUser(String u, String p) {
        username = u;
        password = p;
    }
}

public class TransientDemo {
    public static void main(String[] args) throws Exception {
        SecureUser u = new SecureUser("admin", "1234");
        
        // Write
        ByteArrayOutputStream bos = new ByteArrayOutputStream();
        ObjectOutputStream out = new ObjectOutputStream(bos);
        out.writeObject(u);
        
        // Read
        ObjectInputStream in = new ObjectInputStream(new ByteArrayInputStream(bos.toByteArray()));
        SecureUser res = (SecureUser) in.readObject();
        
        System.out.println("User: " + res.username);
        System.out.println("Pass: " + res.password); // null
    }
}
```

---

## Question 33: Externalizable interface vs Serializable.

### Answer
`Externalizable` gives manual control (`writeExternal`, `readExternal`).

### Runnable Code
```java
package basics;

import java.io.*;

class CustomSer implements Externalizable {
    String data;
    
    public CustomSer() { // No-arg constructor Required
        System.out.println("Constructor called (Externalizable requires this)");
    }
    
    public CustomSer(String d) { data = d; }

    @Override
    public void writeExternal(ObjectOutput out) throws IOException {
        out.writeObject(data.toUpperCase()); // Custom logic (e.g., transform)
    }

    @Override
    public void readExternal(ObjectInput in) throws IOException, ClassNotFoundException {
        data = (String) in.readObject();
    }
}

public class ExternalizableDemo {
    public static void main(String[] args) throws Exception {
        CustomSer c = new CustomSer("hello");
        
        ByteArrayOutputStream bos = new ByteArrayOutputStream();
        new ObjectOutputStream(bos).writeObject(c);
        
        CustomSer res = (CustomSer) new ObjectInputStream( new ByteArrayInputStream(bos.toByteArray())).readObject();
        System.out.println("Result: " + res.data); // HELLO
    }
}
```

---

## Question 34: Byte Stream vs Character Stream.

### Answer
Byte (raw), Char (text/encoding aware).

### Runnable Code
```java
package basics;

import java.io.*;

public class StreamsDemo {
    public static void main(String[] args) throws IOException {
        String content = "Hello World";
        File f = new File("test.txt");
        
        // Character Stream (Writer)
        try (FileWriter fw = new FileWriter(f)) {
            fw.write(content);
        }
        
        // Byte Stream (InputStream)
        try (FileInputStream fis = new FileInputStream(f)) {
            int i;
            System.out.print("Bytes read: ");
            while ((i = fis.read()) != -1) {
                System.out.print((char) i);
            }
        }
        f.delete();
    }
}
```

---

## Question 35: `Scanner` vs `BufferedReader`.

### Answer
Scanner (Parsing), BufferedReader (Fast reading).

### Runnable Code
```java
package basics;

import java.io.*;
import java.util.Scanner;

public class ReaderComparison {
    public static void main(String[] args) throws IOException {
        String input = "10 true Hello";
        
        // Scanner (Parses tokens)
        Scanner sc = new Scanner(input);
        System.out.println("Int: " + sc.nextInt());
        System.out.println("Bool: " + sc.nextBoolean());
        
        // BufferedReader (Reads Lines)
        BufferedReader br = new BufferedReader(new StringReader(input));
        System.out.println("Line: " + br.readLine()); // Reads whole remaining line
    }
}
```

---

## Question 36: Handling `FileNotFoundException` vs `IOException`.

### Answer
Catch specific subclass first.

### Runnable Code
```java
package basics;

import java.io.*;

public class ExceptionOrder {
    public static void main(String[] args) {
        try {
            new FileReader("missing.txt");
        } catch (FileNotFoundException e) {
            System.out.println("Specific: File missing (Subclass of IO)");
        } catch (IOException e) {
            System.out.println("Generic: IO Error");
        }
    }
}
```

---

## Question 37: `File` vs `Path` (NIO.2).

### Answer
`Path` is modern, non-blocking friendly (`java.nio.file`).

### Runnable Code
```java
package basics;

import java.io.File;
import java.nio.file.*;
import java.io.IOException;

public class NIOVsIO {
    public static void main(String[] args) throws IOException {
        // Legacy IO
        File f = new File("legacy.txt");
        f.createNewFile();
        System.out.println("Exists (IO): " + f.exists());
        f.delete();
        
        // NIO.2
        Path p = Paths.get("modern.txt");
        if (Files.notExists(p)) {
            Files.createFile(p);
        }
        System.out.println("Exists (NIO): " + Files.exists(p));
        Files.delete(p);
    }
}
```

---

## Question 38: Breaking Singleton with Serialization.

### Answer
Deserialization creates new instance. Fix: `readResolve`.

### Runnable Code
```java
package basics;

import java.io.*;

class SingletonSer implements Serializable {
    public static final SingletonSer INSTANCE = new SingletonSer();
    private SingletonSer() {}
    
    // Fix: Return existing instance
    protected Object readResolve() {
        return INSTANCE;
    }
}

public class SerializationBreakSingleton {
    public static void main(String[] args) throws Exception {
        SingletonSer s1 = SingletonSer.INSTANCE;
        
        // Serialize
        ByteArrayOutputStream bos = new ByteArrayOutputStream();
        new ObjectOutputStream(bos).writeObject(s1);
        
        // Deserialize
        SingletonSer s2 = (SingletonSer) new ObjectInputStream(new ByteArrayInputStream(bos.toByteArray())).readObject();
        
        System.out.println("Same Instance? " + (s1 == s2)); // true (due to readResolve)
    }
}
```

---

## Question 39: `System.out`, `System.err`, `System.in`.

### Answer
Standard Output, Error, Input streams.

### Runnable Code
```java
package basics;

import java.io.IOException;

public class SysStreams {
    public static void main(String[] args) throws IOException {
        System.out.println("Standard Output (White)");
        System.err.println("Standard Error (Red usually)");
        
        // System.in (Reading byte)
        // int b = System.in.read(); 
    }
}
```

---

## Question 40: Closeable vs AutoCloseable.

### Answer
`AutoCloseable` (Exception), `Closeable` (IOException, extends Auto).

### Runnable Code
```java
package basics;

import java.io.Closeable;
import java.io.IOException;

// Compatible with Try-With-Resources
class Res implements AutoCloseable {
    public void close() throws Exception { // Throws generic Exception
        System.out.println("AutoCloseable closed");
    }
}

// IO Standard
class IORes implements Closeable {
    public void close() throws IOException { // Throws specific IOException
        System.out.println("Closeable closed");
    }
}

public class CloseableDemo {
    public static void main(String[] args) {
        try(Res r = new Res(); IORes io = new IORes()) {
            // Work
        } catch (Exception e) {}
    }
}
```
