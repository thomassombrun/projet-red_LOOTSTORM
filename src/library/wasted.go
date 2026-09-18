package library

import "fmt"

func (c *Character) IsDead() bool {
	if c.CurrentHP <= 0 {
		c.CurrentHP = c.MaxHP / 2
		fmt.Printf("%s est mort... mais ressuscite avec %d PV !\n", c.Name, c.CurrentHP)
		return true
	}
	return false
}
