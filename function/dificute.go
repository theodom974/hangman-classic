package hangman

import (
	"fmt"
	"os"
)

	var fichier *os.File
	var err error

func ChoisirNiveau() string {
	var niveau string

	fmt.Println("Choisissez un niveau de difficulté (facile / difficile / humour / pays) :")
	fmt.Scanln(&niveau)
	//Niveau de dificulté
	switch niveau {
	case "facile":
		fichier, err = os.Open("function/facile.txt")
	case "difficile":
		fichier, err = os.Open("function/difficile.txt")
	case "humour":
		fichier, err = os.Open("function/humour.txt")
		return niveau
	case "pays":
		fichier, err = os.Open("")
	default:
		fmt.Println("Niveau non reconnu. Par défaut, facile.")
	}
	return "facile"
}
	

