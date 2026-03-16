import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

public class Java11_Features {

    public static void main(String[] args) {

        // 1. New String Methods in Java 11
        String text = "   Java 11 Strip   ";

        System.out.println("Original String: [" + text + "]");
        System.out.println("isBlank(): " + text.isBlank());
        System.out.println("strip(): [" + text.strip() + "]");

        String multiline = "A\nB\nC";
        System.out.println("Lines count: " + multiline.lines().count());

        // 2. Java 11 HTTP Client Example
        try {
            HttpClient client = HttpClient.newHttpClient();

            HttpRequest request = HttpRequest.newBuilder()
                    .uri(URI.create("https://api.github.com"))
                    .GET()
                    .build();

            HttpResponse<String> response = client.send(
                    request,
                    HttpResponse.BodyHandlers.ofString());

            System.out.println("\nHTTP Response Status Code: " + response.statusCode());
            System.out.println("Response Body:");
            System.out.println(response.body());

        } catch (Exception e) {
            e.printStackTrace();
        }

        // 3. Java 11 Single File Execution Feature
        System.out.println("\nJava 11 Feature: You can run this file without compiling using:");
        System.out.println("java Java11_Features.java");
    }
}