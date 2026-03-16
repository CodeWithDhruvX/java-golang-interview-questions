public class Java1_1_InnerClasses {
    private int outerValue = 100;

    // Java 1.1 introduced Inner Classes
    class InnerClass {
        void display() {
            System.out.println("Inner class accessing outer field: " + outerValue);
        }
    }

    public static void main(String[] args) {
        Java1_1_InnerClasses outer = new Java1_1_InnerClasses();
        Java1_1_InnerClasses.InnerClass inner = outer.new InnerClass();
        inner.display();

        System.out.println("Java 1.1 features: Inner Classes, JDBC, RMI, Reflection API, and JavaBeans.");
    }
}
