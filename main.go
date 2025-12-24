package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MaminirinaEdwino/compressiondedonnee/module"
)

func main() {
	huffman := flag.Bool("huffman", false, "compresser avec algo Huffman")
	rle := flag.Bool("rle", false, "rle algo")
	bwt := flag.Bool("bwt", false, "bwt algo")
	decodebwt := flag.Bool("decodebwt", false, "decode bwt")
	position := flag.Int("position", 0, "position for rle or bwt string")
	decoderle := flag.Bool("decoderle", false, "decode rle")
	bwtHuffman := flag.Bool("bwthuffman", false, "bwt huffman")
	mot := flag.String("mot", "", "le mot a compressé")
	readfile := flag.Bool("readfile", false, "read compressed file")
	filename := flag.String("file", "", "filename")
	decodeHuffman := flag.Bool("decodehuffman", false, "decode huffman")
	sortie := flag.String()
	flag.Parse()

	file, err := os.Create("compression")
	if err != nil {
		panic(err)
	}

	switch {
	case *readfile:
		if *filename != ""{
			file, err := os.ReadFile(*filename)
			if err != nil {
				panic(err)
			}
			fmt.Println("binary")
			for _, el := range file {
				module.Binary(int(el))
			}
		}
	case *huffman:
		if *mot != "" {
			fmt.Println(*mot)
			// fmt.Println()
			r, f := module.Huffman(*mot)
			module.WriteCompressedIntoFile(r, "sortie", f)
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
			fmt.Println(module.BWT(*mot))
		}
	case *bwtHuffman:
		if *mot != "" {
			res, id := module.BwtHuffman(*mot)
			fmt.Println(res, id)
			file.Write([]byte(res))
		}
	case *decodebwt:
		if *mot != "" && *position != 0 {
			module.DecodeBWT(*mot, *position)
		}
	default:
		module.DecodeBinary("1000")
	}
}
