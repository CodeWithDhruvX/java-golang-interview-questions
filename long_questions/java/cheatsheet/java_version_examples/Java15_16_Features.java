public class Java15_16_Features {
    
    // Java 16: Records became a standard feature (previously preview)
    public record Developer(String name, String language) {}

    public static void main(String[] args) {
        System.out.println("Java 15 & 16 Features Demo");

        // 1. Records (Standardized in 16)
        Developer dev = new Developer("Dhruv", "Java");
        System.out.println("Standardized Record: " + dev);

        // 2. Sealed Classes (Preview in 15/16)
        // (Note: The keyword 'permits' was refined here)
        System.out.println("Sealed Classes (Preview in 15) allowed restricting inheritance.");
        
        // 3. Pattern Matching for instanceof (Standardized in 16)
        Object obj = "Java 16";
        if (obj instanceof String s) {
            System.out.println("Instanceof pattern matching (Standard in 16): " + s.toLowerCase());
        }

        System.out.println("Java 16 also introduced foreign-memory access API (Incubator).");
    }
}
