package module

func BwtHuffman(mot string) (string, int) {
	res, pos := BWT(mot)

	return Huffman(res), pos
}