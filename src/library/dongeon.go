package library

import "fmt"

func PlayDungeonRoom(c *Character, roomNumber int) {

	if c.CurrentHP <= 0 {
		c.Respawn()
		return
	}
	fmt.Printf("Bravo ! Vous avez terminé la salle %d.\n", roomNumber)
	c.SaveGame(roomNumber)
}
