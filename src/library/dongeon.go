package library

import (
	"fmt"
	"math/rand"
	"time"
)

func StartAdventure(c *Character) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println()
	fmt.Println("===== MODE AVENTURE =====")
	fmt.Println("Les salles du donjon sont générées aléatoirement.")
	WaitForEnter()

	for {
		roomNumber := c.LastClearedRoom + 1
		ClearTerminal()
		fmt.Printf("===== SALLE %d =====\n", roomNumber)

		roomCleared := playRandomRoom(c, roomNumber, random)
		if !roomCleared {
			return
		}

		PlayDungeonRoom(c, roomNumber)
		if !continueAdventure() {
			return
		}
	}
}

func playRandomRoom(c *Character, roomNumber int, random *rand.Rand) bool {
	roll := random.Intn(100)
	bossChance := roomNumber * 2
	if bossChance < 5 {
		bossChance = 5
	}
	if bossChance > 40 {
		bossChance = 40
	}
	if roll < bossChance {
		return playBossRoom(c, roomNumber)
	}

	roll -= bossChance
	switch {
	case roll < 35:
		return playGoblinRoom(c, roomNumber)
	case roll < 55:
		return playSlimeRoom(c, roomNumber)
	case roll < 75:
		return playNpcRoom(c)
	default:
		return playChestRoom(c, random)
	}
}

func playGoblinRoom(c *Character, roomNumber int) bool {
	monster := InitGoblinLevel("Gobelin", roomNumber)
	fmt.Printf("Un %s de niveau %d vous attaque !\n", monster.Name, monster.Level)
	return normalFight(c, &monster, roomNumber-1)
}

func playSlimeRoom(c *Character, roomNumber int) bool {
	monster := InitSlimeLevel("Slime", roomNumber)
	fmt.Printf("Un %s de niveau %d vous attaque !\n", monster.Name, monster.Level)
	return normalFight(c, &monster, roomNumber-1)
}

func playBossRoom(c *Character, roomNumber int) bool {
	monster := InitGoblinLevel("Boss gobelin", roomNumber+2)
	monster.MaxHP += 40
	monster.CurrentHP = monster.MaxHP
	monster.Attack += 5
	monster.Initiative += 10
	monster.XPReward += 100
	monster.GoldReward += 50
	fmt.Printf("Un boss de niveau %d apparaît dans la salle !\n", monster.Level)
	return normalFight(c, &monster, roomNumber-1)
}

func playNpcRoom(c *Character) bool {
	fmt.Println("Vous rencontrez un aventurier blessé.")
	fmt.Println("1. Lui demander de vous soigner")
	fmt.Println("2. Continuer votre route")
	fmt.Print("Votre choix : ")

	var choice int
	fmt.Scanln(&choice)
	if choice == 1 {
		healAmount := c.MaxHP / 4
		c.CurrentHP += healAmount
		if c.CurrentHP > c.MaxHP {
			c.CurrentHP = c.MaxHP
		}
		fmt.Printf("L'aventurier vous soigne de %d PV.\n", healAmount)
	} else {
		fmt.Println("Vous quittez la salle sans aide.")
	}
	WaitForEnter()
	return true
}

func playChestRoom(c *Character, random *rand.Rand) bool {
	gold := random.Intn(21) + 10
	c.Gold += gold
	fmt.Printf("Vous trouvez un coffre contenant %d pièces d'or !\n", gold)

	if random.Intn(2) == 0 {
		fmt.Println("Le coffre contient également une Potion de vie.")
		c.AddOrMergeItem("Potion de vie", 1)
	}

	if random.Intn(2) == 0 {
		fmt.Println("Le coffre contient également une Potion de poison.")
		c.AddOrMergeItem("Potion de poison", 1)
	}

	if random.Intn(2) == 0 {
		fmt.Println("Le coffre contient également une Potion de mana.")
		c.AddOrMergeItem("Potion de mana", 1)
	}

	if random.Intn(2) == 0 {
		equipment := []string{
			"Chapeau de l'aventurier",
			"Tunique de l'aventurier",
			"Bottes de l'aventurier",
		}
		foundEquipment := equipment[random.Intn(len(equipment))]
		fmt.Printf("Le coffre contient également : %s.\n", foundEquipment)
		c.AddOrMergeItem(foundEquipment, 1)
	}
	WaitForEnter()
	return true
}

func continueAdventure() bool {
	for {
		ClearTerminal()
		fmt.Println("La salle est terminée.")
		fmt.Println("1. Continuer l'aventure")
		fmt.Println("0. Quitter le donjon")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			return true
		case 0:
			return false
		default:
			fmt.Println("Choix invalide.")
			WaitForEnter()
		}
	}
}

func PlayDungeonRoom(c *Character, roomNumber int) {

	if c.CurrentHP <= 0 {
		c.Respawn()
		return
	}
	fmt.Printf("Bravo ! Vous avez terminé la salle %d.\n", roomNumber)
	c.SaveGame(roomNumber)
}
