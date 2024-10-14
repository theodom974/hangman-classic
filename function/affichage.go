package hangman

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func lirePositions() ([]string, error) {

	contenu, err := ioutil.ReadFile("function/bonhomme.txt")
	if err != nil {
		return nil, err
	}

	positions := strings.Split(string(contenu), "\n\n")
	return positions, nil
}

func afficherPosition(positions []string, tentative int) {
	if tentative >= 0 && tentative < len(positions) {
		fmt.Println(positions[10-tentative])
	}
}
func afficherMotActuel(guesses []rune) {
	fmt.Println("Mot à deviner :", string(guesses))
}

func afficherTentativesRestantes(tentatives int) {
	fmt.Printf("Il vous reste %d tentatives.\n", tentatives)
}

func DemarrerJeu(mot string, tentatives int) {
	motRune := []rune(mot)
	guesses := make([]rune, len(motRune))
	for i := range guesses {
		guesses[i] = '_'
	}

	lettresUtilisees := make(map[rune]bool)

	positions, err := lirePositions()
	if err != nil {
		fmt.Println("Erreur lors de la lecture des positions du bonhomme :", err)
		return
	}

	for tentatives > 0 {
		afficherMotActuel(guesses)
		afficherTentativesRestantes(tentatives)

		var lettre string
		fmt.Print("Proposez une lettre : ")
		fmt.Scanln(&lettre)

		if len(lettre) != 1 {
			fmt.Println("Veuillez entrer une seule lettre.")
			continue
		}

		char := rune(lettre[0])
		if lettresUtilisees[char] {
			fmt.Println("Vous avez déjà utilisé cette lettre.")
			continue
		}

		lettresUtilisees[char] = true
		correct := false
		for i, r := range motRune {
			if r == char {
				guesses[i] = char
				correct = true
			}
		}

		if !correct {
			tentatives--
			afficherPosition(positions, tentatives)
		}

		if strings.Compare(string(guesses), mot) == 0 {
			fmt.Println("Bien joué a toi !!! tu as trouvé le mot :", mot)
			return
		}
	}

	fmt.Println("Désolé, vous avez perdu. Le mot était :", mot)
}
