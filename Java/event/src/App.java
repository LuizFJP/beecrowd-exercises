import java.util.Scanner;

public class App {
    public static void main(String[] args) throws Exception {
        Scanner keyboard = new Scanner(System.in);

        int X = keyboard.nextInt();
        int M = keyboard.nextInt();

        System.out.println(M*X);

        keyboard.close();
    }
}
