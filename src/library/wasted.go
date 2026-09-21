package library

import "fmt"

func (c *Character) IsDead() bool {
	return c.CurrentHP <= 0
}

func (c *Character) Respawn() {
	c.CurrentHP = c.MaxHP / 2
	if c.CurrentHP < 1 {
		c.CurrentHP = 1
	}

	fmt.Println()
	fmt.Println("☠️ Vous avez péri au combat...")
	fmt.Printf("Grâce à votre dernière sauvegarde, vous respawnez au début de la salle %d (avec %d PV).\n",
		c.LastClearedRoom+1,
		c.CurrentHP,
	)
}
