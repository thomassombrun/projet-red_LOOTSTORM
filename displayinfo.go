package main

import "fmt"

func DisplayInfo(c Character) {
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Class: %s\n", c.Class)
	fmt.Printf("Level: %d\n", c.Level)
	fmt.Printf("MaxHP: %d\n", c.MaxHP)
	fmt.Printf("CurrentHP: %d\n", c.CurrentHP)
}

func main() {
	c1.DisplayInfo()
	fmt.Println("---")
	c2.DisplayInfo()
	fmt.Println("---")
	c3.DisplayInfo()
	fmt.Println("---")
}
