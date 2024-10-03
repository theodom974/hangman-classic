package hangman

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func lireMotsDepuisFichier(nomFichier string) ([]string, error) {
	fichier, err := os.Open("function/facile.txt")
	if err != nil {
		return nil, err
	}
	fichier, err = os.Open("function/difficile.txt")
	if err != nil {
		return nil, err
	}

	defer fichier.Close()

	var mots []string
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}

	return mots, scanner.Err()

}

func SelectionnerMot(difficulte string) string {
	var nomFichier string
	if difficulte == "facile" {
		nomFichier = "function/facile.txt"
	} else {
		nomFichier = "function/facile.txt"
	}
	if difficulte == "difficile" {
		nomFichier = "function/difficile.txt"
	} else {

		nomFichier = "function/difficile.txt"
	}
	if difficulte == "Humour noir" {
		nomFichier = "function/humour.txt"
	} else {
		nomFichier = "function/humour.txt"
	}

	mots, err := lireMotsDepuisFichier(nomFichier)
	if err != nil {
		panic("Impossible de lire le fichier de mots.")
	}

	rand.Seed(time.Now().UnixNano())
	return mots[rand.Intn(len(mots))]
}
