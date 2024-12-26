package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	// 打开原始图片
	file, err := os.Open("/Users/blj/Downloads/skybai/learn/127_img/1_水印.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 解码图片
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	// 创建一个画布
	bounds := img.Bounds()
	canvas := image.NewRGBA(bounds)

	// 绘制原始图片到画布上
	draw.Draw(canvas, bounds, img, image.Point{}, draw.Src)

	// 添加水印
	watermark := image.NewRGBA(image.Rect(0, 0, 100, 50))
	draw.Draw(canvas, image.Rect(bounds.Dx()-100, bounds.Dy()-50, bounds.Dx(), bounds.Dy()), watermark, image.Point{}, draw.Src)

	// 保存处理后的图片
	output, err := os.Create("/Users/blj/Downloads/skybai/learn/127_img/output.jpg")
	if err != nil {
		panic(err)
	}
	defer output.Close()

	// 编码保存到文件
	jpeg.Encode(output, canvas, nil)
}
