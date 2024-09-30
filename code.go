package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Liste des mots possibles
var Word }
	"ordinateur", "programmation", "gopher", "clavier", "internet",
	"algorithme", "logiciel", "serveur", "application", "compilateur",
}

// Choisit un mot aléatoire de la liste des mots
func choisirMot() string {
	rand.Seed(time.Now().UnixNano())
	return mots[rand.Intn(len(mots))]
}

// Affiche l'état actuel du mot (les lettres devinées et les tirets)
func afficherMotDevine(mot string, lettresDevinees []string) string {
	motAffiche := ""
	for _, lettre := range mot {
		if contient(lettresDevinees, string(lettre)) {
			motAffiche += string(lettre)
		} else {
			motAffiche += "_"
		}
	}
	return motAffiche
}

// Vérifie si une lettre a déjà été devinée
func contient(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	motADeviner := choisirMot()
	lettresDevinees := []string{}
	nbErreurs := 0
	nbTentativesMax := 10

	fmt.Println("Bienvenue au jeu du pendu !")
	fmt.Printf("Vous avez %d tentatives pour deviner le mot.\n\n", nbTentativesMax)

	for nbErreurs < nbTentativesMax {
		fmt.Println("Mot à deviner :", afficherMotDevine(motADeviner, lettresDevinees))
		fmt.Printf("Vous avez fait %d erreur(s). Tentative restante : %d.\n", nbErreurs, nbTentativesMax-nbErreurs)
		fmt.Print("Devinez une lettre : ")

		var lettre string
		fmt.Scanln(&lettre)
		lettre = strings.ToLower(lettre)

		if len(lettre) != 1 {
			fmt.Println("Veuillez entrer une seule lettre.")
			continue
		}

		if contient(lettresDevinees, lettre) {
			fmt.Println("Vous avez déjà deviné cette lettre.")
			continue
		}

		lettresDevinees = append(lettresDevinees, lettre)

		if strings.Contains(motADeviner, lettre) {
			fmt.Println("Bonne lettre !")
		} else {
			nbErreurs++
			fmt.Println("Mauvaise lettre !")
		}

		if afficherMotDevine(motADeviner, lettresDevinees) == motADeviner {
			fmt.Println("Félicitations ! Vous avez trouvé le mot :", motADeviner)
			return
		}
	}

	fmt.Println("Vous avez perdu ! Le mot était :", motADeviner)
}
