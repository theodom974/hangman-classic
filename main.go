package main

import (
	fn "hangman/function"
	
)

func main() {
	var jeu fn.Jeu
	jeu.ChoisirNiveau()

	jeu.LireMotsDepuisFichier(jeu.Fichier)
	jeu.Mot = jeu.SelectionnerMot()

	
	tentatives := 10
	fn.DemarrerJeu(jeu.Mot, tentatives)
}
