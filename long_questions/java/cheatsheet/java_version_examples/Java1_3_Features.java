import javax.sound.sampled.*;
import java.io.File;

public class Java1_3_Features {
    public static void main(String[] args) {
        // Java 1.3 (2000): HotSpot JVM (Performance) & Java Sound API
        System.out.println("Java 1.3 Features Demo");
        
        // The biggest feature was the HotSpot JVM, which isn't a code feature 
        // but a runtime optimization that significantly boosted Java performance.

        // Java Sound API Example (Conceptual)
        try {
            System.out.println("Java Sound API was introduced to handle audio capture and playback.");
            // Example structure (won't play anything without a real file)
            // AudioInputStream stream = AudioSystem.getAudioInputStream(new File("test.wav"));
            // Clip clip = AudioSystem.getClip();
            // clip.open(stream);
        } catch (Exception e) {
            System.out.println("No sound device or file found.");
        }

        System.out.println("Java 1.3 also introduced JNDI (Java Naming and Directory Interface) as a standard.");
    }
}
