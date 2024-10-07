package hangman

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func (j *Jeu) LireMotsDepuisFichier(fichier *os.File) {

	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		j.WordList = append(j.WordList, scanner.Text())
	}

}

func (j *Jeu) SelectionnerMot() string {
	
	rand.Seed(time.Now().UnixNano())
	return j.WordList[rand.Intn(len(j.WordList))]
}
