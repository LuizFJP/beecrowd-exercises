import java.util.Scanner;

public class App {
    public static void main(String[] args) throws Exception {
         Scanner keyboard = new Scanner(System.in);

         int N = keyboard.nextInt();

        System.out.println(((N+1)*(N+2))/2);

        keyboard.close();
    }
}
