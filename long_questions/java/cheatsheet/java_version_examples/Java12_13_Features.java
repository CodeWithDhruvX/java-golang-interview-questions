public class Java12_13_Features {
    public static void main(String[] args) {
        // Java 12 & 13: Evolutionary features (Previews)
        System.out.println("Java 12 & 13 Features Demo");

        // 1. Switch Expressions (Preview in 12/13, later became standard)
        // Using the 'yield' keyword which was introduced in 13
        String day = "MONDAY";
        int workHours = switch (day) {
            case "MONDAY", "TUESDAY", "WEDNESDAY" -> 8;
            case "SATURDAY", "SUNDAY" -> 0;
            default -> {
                System.out.println("Mid-week processing...");
                yield 6; // yield was introduced in Java 13
            }
        };
        System.out.println("Work hours: " + workHours);

        // 2. Text Blocks (Preview in 13)
        // Solved the problem of multi-line strings
        String multiline = """
                This is a Text Block
                introduced as a preview in Java 13.
                It makes JSON/SQL/HTML strings much cleaner.
                """;
        System.out.println(multiline);
    }
}
