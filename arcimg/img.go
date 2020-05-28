package arcimg

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func init() {
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println("读取字体数据出错")
		log.Fatalln(err)
	}
	font, err = freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println("转换字体样式出错")
		log.Fatalln(err)
	}
}

var font *truetype.Font

const (
	dx = 615 // 图片的大小 宽度
	dy = 212 // 图片的大小 高度
	// fontFile = "FZFSK.TTF"
	fontFile = "t.ttf"
	fontSize = 8   // 字体尺寸
	fontDPI  = 150 // 屏幕每英寸的分辨率
)

func createimg(w io.Writer, info *arcinfo) {
	// 需要保存的文件
	// 新建一个 指定大小的 RGBA位图
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	// 画背景
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			// 设置某个点的颜色，依次是 RGBA
			img.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}
	// 读字体数据

	c := freetype.NewContext()
	c.SetDPI(fontDPI)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.Black)

	pt := freetype.Pt(460, 105) // 字出现的位置
	_, err := c.DrawString(info.Value[0].Avalue.Friends[0].Name, pt)

	pt = freetype.Pt(84, 68)
	if len(getsongname(info.Value[0].Avalue.Friends[0].Recentscore[0].SongID)) > 15 {
		pt = freetype.Pt(66, 68)
	}
	_, err = c.DrawString(getsongname(info.Value[0].Avalue.Friends[0].Recentscore[0].SongID)+"("+info.SongID()+")", pt)

	pt = freetype.Pt(84, 95)
	_, err = c.DrawString(strconv.Itoa(info.Value[0].Avalue.Friends[0].Recentscore[0].Score), pt)

	pt = freetype.Pt(84, 119)
	_, err = c.DrawString(info.atype(), pt)

	pt = freetype.Pt(84, 146)
	_, err = c.DrawString(info.Time(), pt)

	pt = freetype.Pt(268, 68)
	_, err = c.DrawString(info.Pure(), pt)

	pt = freetype.Pt(280, 95)
	_, err = c.DrawString(info.Far(), pt)

	pt = freetype.Pt(270, 119)
	_, err = c.DrawString(info.Lost(), pt)

	pt = freetype.Pt(268, 146)
	_, err = c.DrawString(info.Rating(), pt)

	pt = freetype.Pt(457, 68)
	_, err = c.DrawString(info.PTT(), pt)

	if err != nil {
		log.Println("向图片写字体出错")
		log.Println(err)
		return
	}

	// 以PNG格式保存文件
	err = png.Encode(w, img)
	if err != nil {
		log.Println("生成图片出错")
		log.Println(err)
	}
}
