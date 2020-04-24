package top.xmdhs.arcimg;

import javax.imageio.ImageIO;
import javax.imageio.stream.ImageOutputStream;
import java.awt.*;
import java.awt.image.BufferedImage;
import java.io.*;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class CreatImg {
    public static final Lock lock = new ReentrantLock();
    public static final Lock alock = new ReentrantLock();
    public static InputStream image;
    public static long time = 0;

    public static InputStream creatImg() {
        int imageWidth = 615;//图片的宽度
        int imageHeight = 212;//图片的高度
        alock.lock();
        try {
            if (System.currentTimeMillis() - time < 30000) {
                time = System.currentTimeMillis();
                return getInputStream();
            }
        } finally {
            alock.unlock();
        }
        BufferedImage image = new BufferedImage(imageWidth, imageHeight, BufferedImage.TYPE_INT_RGB);
        Graphics graphics = image.getGraphics();
        lock.lock();
        try {
            if (System.currentTimeMillis() - a.time > 600000) {
                a.json = json2.json2class(http.https());
                a.time = System.currentTimeMillis();
            }
        } finally {
            lock.unlock();
        }
        try {
            Font font = new Font("sans-serif", Font.PLAIN, 16);
            graphics.setFont(font);
            graphics.fillRect(0, 0, imageWidth, imageHeight);
            graphics.setColor(new Color(0, 0, 0));//设置黑色字体,同样可以graphics.setColor(Color.black);
            graphics.drawString(text.song(), 84, 68);
            graphics.drawString(a.json.value.get(0).value.friends.get(0).recent_score.get(0).score, 84, 95);
            graphics.drawString(text.type(), 84, 119);
            graphics.drawString(text.time(), 84, 146);
            graphics.drawString(text.pure(), 268, 68);
            graphics.drawString(text.far(), 280, 95);
            graphics.drawString(text.lost(), 270, 119);
            graphics.drawString(text.rating(), 268, 146);
            graphics.drawString(text.ptt(), 457, 68);
            graphics.drawString("xmdhs", 460, 105);
            return getImageStream(image);
        } catch (Exception ex) {
            ex.printStackTrace();
            return null;
        }


    }
    private static ByteArrayOutputStream bs = null;

    public static InputStream getImageStream(BufferedImage bimage) {
        ImageOutputStream imOut;
        bs = new ByteArrayOutputStream();
        try {
            imOut = ImageIO.createImageOutputStream(bs);
            ImageIO.write(bimage, "png", imOut);
            bs.flush();
        } catch (IOException e) {
            e.printStackTrace();
        }
        return getInputStream();
    }
    public static InputStream getInputStream() {
        return new ByteArrayInputStream(bs.toByteArray());
    }
}

