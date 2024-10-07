package hangman

import "os"

type Jeu struct {
	Fichier *os.File
	WordList []string

	Mot           string
	Tentatives    int
	LettresDevines map[rune]bool
}
