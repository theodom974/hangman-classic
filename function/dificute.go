package hangman

import (
	"fmt"
	"os"
)

var err error

func (j *Jeu) ChoisirNiveau() {
	var niveau string

	fmt.Println("Choisissez un niveau de difficulté (facile / difficile / marque / pays) :")
	fmt.Scanln(&niveau)
	//Niveau de dificulté
	switch niveau {
	case "facile":
		j.Fichier, err = os.Open("function/facile.txt")
	case "difficile":
		j.Fichier, err = os.Open("function/difficile.txt")
	case "marque":
		j.Fichier, err = os.Open("function/marque.txt")
	case "pays":
		j.Fichier, err = os.Open("function/pays.txt")
	default:
		fmt.Println("Niveau non reconnu. Par défaut, facile.")
		j.ChoisirNiveau()
	}
}
