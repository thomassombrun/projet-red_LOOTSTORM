package library

import "fmt"

func (c Character) DisplayInfo() {
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Class: %s\n", c.Class)
	fmt.Printf("Level: %d\n", c.Level)
	fmt.Printf("MaxHP: %d\n", c.MaxHP)
	fmt.Printf("CurrentHP: %d\n", c.CurrentHP)
}
