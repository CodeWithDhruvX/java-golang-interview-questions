import com.sun.net.httpserver.HttpServer;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpExchange;
import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;

public class Java18_WebServer {
    public static void main(String[] args) throws IOException {
        System.out.println("Java 18 Features Demo");
        
        // Java 18: Simple Web Server (jwebserver)
        // This code demonstrates the internal API used by the jwebserver tool
        
        int port = 8080;
        HttpServer server = HttpServer.create(new InetSocketAddress(port), 0);
        server.createContext("/", new MyHandler());
        server.setExecutor(null); // creates a default executor
        
        System.out.println("Starting simple web server on port " + port);
        System.out.println("Hit Ctrl+C to stop.");
        
        // Note: For demonstration purposes, we start it briefly
        server.start();
        
        System.out.println("Server started effectively (Sample handle provided).");
        
        // Immediately stopping for the sake of the demo (in real use, keep it running)
        server.stop(1);

        System.out.println("Java 18 also made UTF-8 the default charset across all platforms.");
    }

    static class MyHandler implements HttpHandler {
        @Override
        public void handle(HttpExchange t) throws IOException {
            String response = "Hello from Java 18 Simple Web Server!";
            t.sendResponseHeaders(200, response.length());
            OutputStream os = t.getResponseBody();
            os.write(response.getBytes());
            os.close();
        }
    }
}
