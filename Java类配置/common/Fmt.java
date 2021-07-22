package common;

// 这个类用于打印字符串
public class Fmt {
    private static native void print0(String str);

    public static void println(String str){
        print0(str);
    }
}
