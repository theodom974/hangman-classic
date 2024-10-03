package hangman

type Jeu struct {
	Mot           string
	Tentatives    int
	LettresDevines map[rune]bool
}
