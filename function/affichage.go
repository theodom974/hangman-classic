package hangman

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
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

// Initialise le mot à deviner et révèle aléatoirement quelques lettres
func initialiseMot(mot string, nbLettresRevelees int) []rune {
	longueurMot := len(mot)
	revelation := make([]rune, longueurMot)

	// Initialise avec des tirets pour chaque lettre
	for i := range revelation {
		revelation[i] = '_'
	}

	// Révèle aléatoirement certaines lettres
	rand.Seed(time.Now().UnixNano())
	if nbLettresRevelees > longueurMot {
		nbLettresRevelees = longueurMot // Ajuste au cas où il y aurait trop de lettres demandées
	}
	indicesReveles := rand.Perm(longueurMot)[:nbLettresRevelees]
	for _, i := range indicesReveles {
		revelation[i] = rune(mot[i])
	}

	return revelation
}

func DemarrerJeu(mot string, tentatives int) {
	motRune := []rune(mot)
	guesses := initialiseMot(mot, 2) // Révèle 2 lettres au hasard
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
			fmt.Println("Bien joué à toi !!! Tu as trouvé le mot :", mot)
			return
		}
	}

	fmt.Println("Désolé, vous avez perdu. Le mot était :", mot)
}
