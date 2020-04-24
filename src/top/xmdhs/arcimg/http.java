package top.xmdhs.arcimg;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.nio.charset.StandardCharsets;

public class http {
    public static String https(){
        try {
            URL url = new URL("https://arcapi.lowiro.com/11/compose/aggregate?calls=%5B%7B%20%22endpoint%22%3A%20%22%2Fuser%2Fme%22%2C%20%22id%22%3A%200%20%7D%5D");
            HttpURLConnection connection = (HttpURLConnection) url.openConnection();
            connection.setRequestMethod("GET");
            connection.setUseCaches(false);
            connection.setConnectTimeout(5000);
            connection.setRequestProperty("Accept-Encoding", "identity");
            connection.setRequestProperty("Content-Type", "application/x-www-form-urlencoded; charset=utf-8");
            connection.setRequestProperty("Authorization", "");//填入抓包获取的
            connection.setRequestProperty("Accept-Encoding", "identity");
            connection.setRequestProperty("Platform", "android");
            connection.setRequestProperty("AppVersion", "2.6.1c");
            connection.setRequestProperty("i", "2550615");
            connection.setRequestProperty("User-Agent", "Dalvik/2.1.0 (Linux; U; Android 10; GM1900 Build/QKQ1.190716.003)");
            connection.setRequestProperty("Host", "arcapi.lowiro.com");
            connection.setRequestProperty("Connection", "Keep-Alive");
            BufferedReader in = new BufferedReader(new InputStreamReader(connection.getInputStream(), StandardCharsets.UTF_8));
            StringBuilder json = new StringBuilder();
            String current;
            while ((current = in.readLine()) != null) {
                json.append(current);
            }
            return json.toString();
        } catch (Exception e) {
            e.printStackTrace();
            return "";
        }

    }
}
