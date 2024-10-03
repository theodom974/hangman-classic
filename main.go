package main

import (
	fn"hangman/function"
	
)

func main() {
	niveau := fn.ChoisirNiveau()

	
	mot := fn.SelectionnerMot(niveau)

	
	tentatives := 10
	fn.DemarrerJeu(mot, tentatives)
}
