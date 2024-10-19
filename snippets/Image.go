package snippets

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func findModel(model color.Model) {
	switch model {
	case color.RGBAModel:
		fmt.Println("The image uses the RGBA color model.")
	case color.NRGBAModel:
		fmt.Println("The image uses the NRGBA color model.")
	case color.GrayModel:
		fmt.Println("The image is grayscale.")
	case color.CMYKModel:
		fmt.Println("The image uses the CMYK color model.")
	case color.YCbCrModel:
		fmt.Println("The image uses the YCbCr color model.")
	default:
		fmt.Printf("Unknown color model: %T\n", model)
	}

}
func ConvertRed(img image.Image) image.Image {

	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			fmt.Println("color = ", img.At(i, j))
			originalColor := img.At(i, j)

			// returns uint32
			r, _, _, a := originalColor.RGBA()

			r8 := uint8(r >> 8)
			a8 := uint8(a >> 8)

			newImg.Set(i, j, color.RGBA{R: r8, G: 0, B: 0, A: a8})
		}

	}
	return newImg
}

func ConvertGreen(img image.Image) image.Image {

	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			fmt.Println("color = ", img.At(i, j))
			originalColor := img.At(i, j)

			// returns uint32
			_, g, _, a := originalColor.RGBA()

			g8 := uint8(g >> 8)
			a8 := uint8(a >> 8)

			newImg.Set(i, j, color.RGBA{R: 0, G: g8, B: 0, A: a8})
		}

	}
	return newImg
}

func ConvertBlue(img image.Image) image.Image {

	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			fmt.Println("color = ", img.At(i, j))
			originalColor := img.At(i, j)

			// returns uint32
			_, _, b, a := originalColor.RGBA()

			b8 := uint8(b >> 8)
			a8 := uint8(a >> 8)

			newImg.Set(i, j, color.RGBA{R: 0, G: 0, B: b8, A: a8})
		}

	}
	return newImg
}
func saveImageWithColor(filename string, img image.Image) error {
	// Create the output file
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Encode the image to the file
	err = jpeg.Encode(outFile, img, nil)
	if err != nil {
		return err
	}

	return nil
}

// This is the dynamically generated function for your snippet
func (s Snip) ImageMain() {
	reader, err := os.Open("assets/sample_jpeg.jpeg")
	if err != nil {
		log.Fatal("img fetch err ", err)
	}
	defer reader.Close()

	img, err := jpeg.Decode(reader)
	if err != nil {
		log.Fatalf("Failed to decode image: %v", err)
	}

	// gets the dimensions of the image
	fmt.Println("img bounds ", img.Bounds())
	fmt.Printf("img color model %T\n", img.ColorModel())

	redImg := ConvertRed(img)
	greenImg := ConvertGreen(img)
	blueImg := ConvertBlue(img)

    saveImageWithColor("red.jpeg",redImg)
    saveImageWithColor("green.jpeg",greenImg)
    saveImageWithColor("blue.jpeg",blueImg)
}
