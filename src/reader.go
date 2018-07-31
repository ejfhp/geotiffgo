package main

import (
	"bufio"
	"flag"
	"fmt"
	gtiff "github.com/google/tiff"
	"golang.org/x/image/tiff"
	"os"
)

func main() {
	fmt.Printf("Reading GeoTiff\n")
	fileName := flag.String("tiff", "", "Path of the TIFF")
	flag.Parse()

	fmt.Printf("TIFF file is %v\n", *fileName)
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("Cannot open file due to %v\n", err)
		os.Exit(-1)
	}
	reader := bufio.NewReader(file)
	tiffImg, err := tiff.Decode(reader)
	if err != nil {
		fmt.Printf("Cannot read image due to %v", err)
	}
	tiffConfig, err := tiff.DecodeConfig(reader)
	if err != nil {
		fmt.Printf("Cannot read image configuration due to %v", err)
	}
	fmt.Printf("Image bounds are: %s\n", tiffImg.Bounds().String())
	fmt.Printf("Image config is: %s\n", tiffConfig.ColorModel)

	tags := gtiff.ListTagSpaceNames()
	fmt.Printf("Number of tags: %d\n", len(tags))
	for tn := range tags {
		fmt.Printf("Tag name: %v\n", tn)
	}
}
