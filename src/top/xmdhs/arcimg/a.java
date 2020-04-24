package top.xmdhs.arcimg;


import java.io.*;
import java.net.ServerSocket;
import java.net.Socket;
import java.nio.charset.StandardCharsets;
import java.util.concurrent.*;

public class a {
    public static long time = 0;
    public static Json json;
    public static ConcurrentHashMap<String, Integer> hashMap = new ConcurrentHashMap<>();

    public static void main(String[] args) throws IOException {
        ExecutorService cachedThreadPool = Executors.newCachedThreadPool();
        ScheduledThreadPoolExecutor exec = new ScheduledThreadPoolExecutor(1);
        exec.scheduleAtFixedRate(() -> {
            for (String i : hashMap.keySet()){
                if(hashMap.get(i) <= 0){
                    hashMap.remove(i);
                }else {
                hashMap.put(i,hashMap.get(i)-1);
                }
            }
        }, 1000,6000,TimeUnit.MILLISECONDS);
        ServerSocket ss = new ServerSocket(8080);
        System.out.println("server is running...");
        //noinspection InfiniteLoopStatement
        for (; ; ) {
            Socket sock = ss.accept();
            System.out.println("connected from " + sock.getRemoteSocketAddress());
            cachedThreadPool.execute(new Handler(sock));
        }
    }

}
class Handler extends Thread {
    Socket sock;

    public Handler(Socket sock) {
        this.sock = sock;
    }

    public void run() {
        try (InputStream input = this.sock.getInputStream()) {
            try (OutputStream output = this.sock.getOutputStream()) {
                handle(input, output);
            }
        } catch (Exception e) {
            try {
                this.sock.close();
            } catch (IOException ignored) {
            }
            System.out.println("client disconnected.");
        }
    }

    private void handle(InputStream input, OutputStream output) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(input, StandardCharsets.UTF_8));
        BufferedWriter writer = new BufferedWriter(new OutputStreamWriter(output, StandardCharsets.UTF_8));
        boolean requestOk = true;
        String first = reader.readLine();
        System.out.println(first);
        if (first.startsWith("GET /favicon.ico HTTP/1.")) {
            requestOk = false;
        }
        for (; ; ) {
            String header = reader.readLine();
            if (header.isEmpty()) { // 读取到空行时, HTTP Header读取完毕
                break;
            }
            if (header.contains("X-Forwarded-For")) {
                System.out.println(header);
                if(a.hashMap.containsKey(header)){
                    a.hashMap.put(header,a.hashMap.get(header) + 1);
                }else {
                    a.hashMap.put(header,1);
                }
                if(a.hashMap.get(header) > 10){
                    a.hashMap.put(header,20);
                    requestOk = false;
                }
            }
        }
        System.out.println(requestOk ? "Response OK" : "Response Error");
        if (!requestOk) {
            // 发送错误响应:
            writer.write("HTTP/1.1 404 Not Found\n");
            writer.write("content-length: 0\n");
            writer.write("server: xmdhs\n\n");
        } else {
            InputStream img = CreatImg.creatImg();
            if (img != null) {
                int i = img.available();
                write(output, "HTTP/1.1 200 OK\n");
                write(output, "server: xmdhs\n");
                write(output, "content-type: image/png\n");
                write(output, "Cache-Control: max-age=60\n");
                write(output, "content-length: " + i);
                write(output, "\n\n");
                copyAllBytes(img, output);
            }
        }
        writer.flush();
    }

    void write(OutputStream output, String s) throws IOException {
        output.write(s.getBytes());
    }

    void copyAllBytes(InputStream input, OutputStream output) throws IOException {
        byte[] buf = new byte[4096];
        int n;
        while ((n = input.read(buf)) != -1) {
            output.write(buf, 0, n);
        }
    }
}




