import common.Fmt;
import common.Thread;

class Lucky extends Thread {

    public void run() {
        for (int i = 0; i <100; i++) {
            Fmt.println("kao");
        }
    }

    private static int max =2;
    public static void main(String[] args) {
        String a = "123";
        String b = "456";

        String c = a + b;

        Lucky lucky=new Lucky();
        lucky.start();

        //Fmt.println(max+"");
        for (int i = 0; i <100; i++) {
            Fmt.println("kao");
        }
    }
}
