package common;

public class Thread {
    // 开启一个新的线程
    private native  void start0();

    // 创建一个子类，继承Thread类，然后重写run方法。
    public void run(){}

    public void start(){
        start0();
    }
}
