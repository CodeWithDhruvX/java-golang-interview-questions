import java.io.Console;
import javax.script.*;

public class Java6_Features {
    public static void main(String[] args) throws ScriptException {
        // Java 6: Scripting Support, JDBC 4.0, Compiler API, Console API
        System.out.println("Java 6 Features Demo");

        // 1. Scripting Engine support (JS integration)
        ScriptEngineManager manager = new ScriptEngineManager();
        ScriptEngine engine = manager.getEngineByName("JavaScript");
        if (engine != null) {
            System.out.print("JavaScript Execution: ");
            engine.eval("print('Hello from JavaScript inside Java 6!')");
        }

        // 2. Console API
        Console console = System.console();
        if (console != null) {
            // Note: Console only works when run from a real terminal, not IDE output
            System.out.println("Console API is available for secure password entry.");
        } else {
            System.out.println("No system console available (common in IDEs).");
        }

        System.out.println("Java 6 also introduced JDBC 4.0 (Automatic driver loading) and Pluggable Annotations.");
    }
}
