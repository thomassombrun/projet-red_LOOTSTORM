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

// SaveGameSilent writes the save file without printing to stdout, for callers
// (like the graphical UI) that autosave frequently and don't want console spam.
func (c *Character) SaveGameSilent() error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("save.json", data, 0644)
}

// HasSaveGame reports whether a save file exists to continue from.
func HasSaveGame() bool {
	_, err := os.Stat("save.json")
	return err == nil
}

// LoadGame reads the save file and returns the saved character.
func LoadGame() (Character, error) {
	var c Character
	data, err := os.ReadFile("save.json")
	if err != nil {
		return c, err
	}
	if err := json.Unmarshal(data, &c); err != nil {
		return c, err
	}
	return c, nil
}
