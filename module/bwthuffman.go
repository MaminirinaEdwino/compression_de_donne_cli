package module

func BwtHuffman(mot string) (string, int) {
	res, pos := BWT(mot)
	hr, _ := Huffman(res)
	return hr, pos
}