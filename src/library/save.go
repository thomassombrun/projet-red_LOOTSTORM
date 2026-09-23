package library

import (
	"encoding/json"
	"fmt"
	"os"
)

func (c *Character) SaveGame(roomNumber int) {
	c.LastClearedRoom = roomNumber

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println("Erreur lors de la préparation de la sauvegarde :", err)
		return
	}
	err = os.WriteFile("save.json", data, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier de sauvegarde :", err)
		return
	}
	fmt.Println()
	fmt.Printf("💾 Partie sauvegardée avec succès ! (Checkpoint : Salle %d validée)\n", roomNumber)
}
