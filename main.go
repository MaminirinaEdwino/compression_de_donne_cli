package main

import (
	"flag"
	"fmt"

	"github.com/MaminirinaEdwino/compressiondedonnee/module"
)

func main() {
	huffman := flag.Bool("huffman", false, "compresser avec algo Huffman")
	rle := flag.Bool("rle", false, "rle algo")
	bwt := flag.Bool("bwt", false, "bwt algo")
	decodebwt := flag.Bool("decodebwt", false, "decode bwt")
	position := flag.Int("position", 0, "position for rle or bwt string")
	decoderle := flag.Bool("decoderle", false, "decode rle")

	mot := flag.String("mot", "", "le mot a compressé")
	flag.Parse()
	switch {
	case *huffman:
		if *mot != "" {
			fmt.Println(*mot)
			module.Huffman(*mot)
		}
	case *rle:
		if *mot != "" {
			module.RLE(*mot)
		}
	case *decoderle:
		if *mot != "" {
			module.DecodeRLE(*mot)
		}
	case *bwt:
		if *mot != "" {
			module.BWT(*mot)
		}
	case *decodebwt:
		if *mot != "" && *position != 0 {
			module.DecodeBWT(*mot, *position)
		}
	}
}
