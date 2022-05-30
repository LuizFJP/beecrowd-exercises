import java.util.Scanner;

public class App {
    public static void main(String[] args) throws Exception {
      Scanner keyboard = new Scanner(System.in);

      int D = keyboard.nextInt();
      int count = 0;

      if (D <= 800 ) count += 1;
      if (800 < D && D <= 1400 ) count += 2;
      if (1400 < D && D <= 2000 ) count += 3;

      System.out.println(count);

      keyboard.close();
    }
}
