package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	_ "image/png"
	"os"
	"time"
)

func main() {
	err := removeLogo("/Users/blj/Downloads/skybai/learn/127_img/img_720.png", "/Users/blj/Downloads/skybai/learn/127_img/output_720.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = removeLogo("/Users/blj/Downloads/skybai/learn/127_img/img_1080.jpg", "/Users/blj/Downloads/skybai/learn/127_img/output_1080.jpg")
	fmt.Println(err)
}

func removeLogo(objectName, out string) error {

	// 打开图像文件
	file, err := os.Open(objectName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 解码图像
	img, _, err := image.Decode(file)
	if err != nil {
		return errors.New("removeLogo image.Decode err: " + err.Error())
	}

	// 创建一个画布
	bounds := img.Bounds()
	canvas := image.NewRGBA(bounds)

	// 绘制原始图像到画布上
	draw.Draw(canvas, bounds, img, image.Point{}, draw.Src)
	fmt.Printf("dx:%d,dy:%d\n", bounds.Dx(), bounds.Dy())

	x1 := 0
	y1 := 0
	watermarkY := 60

	// 水印开始位置
	startX := 24

	// 如果图片的宽大于等于1280,高大于等于720 水印宽为400,高为100 设置水印大小
	if bounds.Dx() >= 1280 && bounds.Dy() >= 720 {
		x1 = 400
		y1 = 100
	} else {
		// 如果图片的宽小于1280,高小于720 水印宽为200,高为50
		x1 = 200
		y1 = 50
	}

	// 如果是大于720 图宽是1920 高是1080
	if bounds.Dx() > 1280 && bounds.Dy() > 720 {
		startX = 24     // 从左开始向右
		watermarkY = 60 // 水印高度
	} else if bounds.Dx() <= 1280 && bounds.Dy() <= 720 {
		startX = 15     // 从左开始向右
		watermarkY = 47 // 水印高度
	}

	// 水印结束位置
	endX := x1 - startX

	stime := time.Now()
	watermark := image.NewRGBA(image.Rect(0, 0, x1, y1))
	gray := color.RGBA{R: 200, G: 200, B: 200, A: 255} // 设置颜色
	for x := 0; x < x1; x++ {
		for y := 0; y < y1; y++ {
			watermark.Set(x, y, gray)
		}
	}

	etime := time.Now()
	fmt.Printf("create watermark time: %v\n", etime.Sub(stime))

	//stime = time.Now()
	draw.Draw(canvas, image.Rect(startX, bounds.Dy()-watermarkY-y1, endX, bounds.Dy()-watermarkY), watermark, image.Point{}, draw.Over)
	etime = time.Now()
	fmt.Printf("draw watermark time: %v\n", etime.Sub(stime))

	// 5. 创建新的输出文件
	outputFile, err := os.Create(out)
	if err != nil {
		return errors.New("removeLogo os.Create output file err: " + err.Error())
	}

	defer func() {
		outputFile.Close()
		//os.Remove(outputFile.Name())
	}()

	// 编码并保存到文件
	err = jpeg.Encode(outputFile, canvas, nil)
	if err != nil {
		return errors.New("removeLogo jpeg.Encode err" + err.Error())
	}
	return nil
}
