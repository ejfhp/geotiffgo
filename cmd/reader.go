package main

import (
	"bufio"
	"flag"
	"fmt"
	gtiff "github.com/google/tiff"
	// "golang.org/x/image/tiff"
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

	//GOLANG TIFF
	// tiffImg, err := tiff.Decode(reader)
	// if err != nil {
	// 	fmt.Printf("Cannot read image due to %v", err)
	// }
	// tiffConfig, err := tiff.DecodeConfig(reader)
	// if err != nil {
	// 	fmt.Printf("Cannot read image configuration due to %v", err)
	// }
	// fmt.Printf("Image bounds are: %s\n", tiffImg.Bounds().String())
	// fmt.Printf("Image config is: %s\n", tiffConfig.ColorModel)

	// tags := gtiff.ListTagSpaceNames()
	// fmt.Printf("Number of tags: %d\n", len(tags))
	// for tn := range tags {
	// 	fmt.Printf("Tag name: %v\n", tn)
	// }

	// reader.Reset(reader)
	// fmt.Printf("Reader reset, now reading with gtiff\n")
	//GTIFF
	readAtReadSeeker := gtiff.NewReadAtReadSeeker(reader)

	tiffImage, err := gtiff.Parse(readAtReadSeeker, gtiff.NewTagSpace("GeoTIFF"), gtiff.NewFieldTypeSpace("Geotiff"))
	if err != nil {
		fmt.Printf("Cannot parse tiff due to %v\n", err)
		os.Exit(-1)
	}
	ifds := tiffImage.IFDs()
	fmt.Printf("Number of pages found: %v\n", len(ifds))
	for idp, page := range ifds {
		fmt.Printf(" Scanning page number %v\n", idp)
		fields := page.Fields()
		tagIDs := make([]uint16, len(fields), len(fields))
		fmt.Printf("  Number of fields found: %v\n", len(fields))
		for idf, field := range fields {
			tagIDs[idf] = field.Tag().ID()
			fmt.Printf("  Scanning field number %v\n", idf)
			fmt.Printf("   Field count:              %d\n", field.Count())
			fmt.Printf("   Field offset:             %d\n", field.Offset())
			fmt.Printf("   Field tag ID:             %d\n", field.Tag().ID())
			fmt.Printf("   Field tag name:           %s\n", field.Tag().Name())
			fmt.Printf("   Field type ID:            %d\n", field.Type().ID())
			fmt.Printf("   Field type name:          %s\n", field.Type().Name())
			fmt.Printf("   Field type type:          %s\n", field.Type().ReflectType().Name())
			fmt.Printf("   Field type signed:        %t\n", field.Type().Signed())
			fmt.Printf("   Field type size:          %d\n", field.Type().Size())
			fmt.Printf("   Field value order:        %s\n", field.Value().Order().String())
			fmt.Printf("   Field value byte:         %s\n", string(field.Value().Bytes()))
		}
		fmt.Printf("Tag IDs: %v\n", tagIDs)
	}
}
