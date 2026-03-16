import java.util.logging.*;
import java.io.*;

public class Java1_4_Features {
    public static void main(String[] args) {
        // Java 1.4: Assert keyword, Logging API, NIO (New I/O)
        
        // Assert keyword (Enable with -ea)
        int testValue = 10;
        assert testValue > 0 : "Value should be positive";
        
        // Logging API
        Logger logger = Logger.getLogger(Java1_4_Features.class.getName());
        logger.info("This is a Java 1.4 Standard Logging message.");

        System.out.println("Java 1.4 features: assertions (run with -ea), Logging API, NIO, and Exception Chaining.");
    }
}
