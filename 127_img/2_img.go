package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	// 打开图像文件
	file, err := os.Open("/Users/blj/Downloads/skybai/learn/127_img/img.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 解码图像
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	// 创建一个画布
	bounds := img.Bounds()
	canvas := image.NewRGBA(bounds)

	// 绘制原始图像到画布上
	draw.Draw(canvas, bounds, img, image.Point{}, draw.Src)

	// 图片长度的一半

	// 创建一个 100x25 大小的浅灰色水印图像
	watermark := image.NewRGBA(image.Rect(0, 0, 200, 50))
	gray := color.RGBA{R: 200, G: 200, B: 200, A: 255}
	for x := 0; x < 200; x++ {
		for y := 0; y < 50; y++ {
			watermark.Set(x, y, gray)
		}
	}

	// 添加水印到画布的左下角
	watermarkPosition := image.Pt(0, bounds.Dy()-100) // 水印左下角的点
	draw.Draw(canvas, image.Rect(watermarkPosition.X, watermarkPosition.Y, watermarkPosition.X+200, watermarkPosition.Y+75), watermark, image.Point{}, draw.Over)

	// 保存处理后的图像
	output, err := os.Create("/Users/blj/Downloads/skybai/learn/127_img/output.jpg")
	if err != nil {
		panic(err)
	}
	defer output.Close()

	// 编码并保存到文件
	jpeg.Encode(output, canvas, nil)
}

// 0 表示水印的横坐标（即水平位置），放在图像的最左边。
//bounds.Dy()-50 表示水印的纵坐标（即垂直位置），bounds.Dy() 是图像的高度，减去 50 得到水印的顶部应该放置的位置，这样水印的底部与图像底部对齐。

// image.Rect(watermarkPosition.X, watermarkPosition.Y, watermarkPosition.X+150, watermarkPosition.Y+50) 定义了在画布上绘制水印的矩形区域。这个区域从 watermarkPosition.X 和 watermarkPosition.Y（即水印的左上角）开始，到 watermarkPosition.X+150 和 watermarkPosition.Y+50（即水印的右下角）结束。
//watermarkPosition.X 是水印的左上角横坐标（0）。
//watermarkPosition.Y 是水印的左上角纵坐标（bounds.Dy()-50）。
//watermarkPosition.X+150 是水印的右下角横坐标（150，水印宽度）。
//watermarkPosition.Y+50 是水印的右下角纵坐标（bounds.Dy()，水印高度加上水印顶部位置等于图像底部）。
