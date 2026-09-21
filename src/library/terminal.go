package library

import "fmt"

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func WaitForEnter() {
	fmt.Print("\nAppuyez sur Entrée pour continuer...")
	fmt.Scanln()
}
