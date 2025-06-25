package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math/rand"
	"os"
	"strconv"
)

func main(){
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Enter args <size> <output_name>")
		return
	}
	img_size, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Enter a valid image size:\nnois_gen <image_size> <output_name>")
		return
	}
	if img_size < 8 || img_size > 2160 {
		fmt.Println("Image size must be in the range [8,2160]")
		return
	}
	noise_image := image.NewRGBA(image.Rect(0,0,img_size, img_size))
	for i:=0; i<img_size; i++ {
		for j:=0; j<img_size; j++ {
			r := uint8(rand.Intn(256))
			g := uint8(rand.Intn(256))
			b := uint8(rand.Intn(256))
			noise_image.Set(i, j, color.RGBA{R:r, G:g, B:b, A:255})
		}
	}
	jpg_file, err := os.Create(args[2]+".jpg")
	if err != nil {
		fmt.Println("Error creating the jpg.")
		return
	}
	defer jpg_file.Close()
	jpeg.Encode(jpg_file, noise_image, &jpeg.Options{Quality: 100})
}
