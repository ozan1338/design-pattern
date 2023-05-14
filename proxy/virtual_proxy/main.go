package main

import "fmt"

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading image from ", filename)
	return &Bitmap{filename}
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing Image", b.filename)
}

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func (this *LazyBitmap) Draw() {
	if this.bitmap == nil {
		this.bitmap = NewBitmap(this.filename)
	}
	this.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{
		filename,
		nil,
	}
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}

func main() {
	bmp := NewLazyBitmap("/demo.png")
	DrawImage(bmp)
}
