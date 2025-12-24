package module

import (
	"fmt"
	"os"
	"strconv"
)

func WriteCompressedIntoFile(huffmanBits string, filename string){

	file, err := os.Create(filename+".huff")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var bytes []byte

	// 2. Parcourir par blocs de 8
	for i := 0; i < len(huffmanBits); i += 8 {
		end := i + 8
		padding := 0

		if end > len(huffmanBits) {
			padding = end - len(huffmanBits)
			end = len(huffmanBits)
		}

		bitString := huffmanBits[i:end]

		// Ajouter des '0' à la fin si nécessaire pour compléter l'octet
		for j := 0; j < padding; j++ {
			bitString += "0"
		}

		// Convertir la chaîne "0101..." en un nombre (uint8)
		val, _ := strconv.ParseUint(bitString, 2, 8)
		bytes = append(bytes, byte(val))
	}

	// 3. Écrire les octets dans le fichier
	_, err = file.Write(bytes)
	if err != nil {
		fmt.Println("Erreur écriture:", err)
	}

	fmt.Printf("Fichier écrit : %d octets générés pour %d bits.\n", len(bytes), len(huffmanBits))
}