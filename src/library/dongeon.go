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
		floor := dungeonFloor(roomNumber)
		ClearTerminal()
		fmt.Printf("===== ÉTAGE %d - SALLE %d =====\n", floor, roomNumber)

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

func dungeonFloor(roomNumber int) int {
	if roomNumber < 1 {
		return 1
	}
	return (roomNumber-1)/10 + 1
}

func playRandomRoom(c *Character, roomNumber int, random *rand.Rand) bool {
	if roomNumber%10 == 0 {
		return playBossRoom(c, roomNumber)
	}

	roll := random.Intn(100)
	bossChance := 10 + roomNumber/2
	if bossChance < 8 {
		bossChance = 8
	}
	if bossChance > 35 {
		bossChance = 35
	}
	if roll < bossChance {
		return playBossRoom(c, roomNumber)
	}

	roll = random.Intn(100 - bossChance)
	switch {
	case roll < 14:
		return playGoblinRoom(c, roomNumber)
	case roll < 26:
		return playSlimeRoom(c, roomNumber)
	case roll < 36:
		return playGhostRoom(c, roomNumber)
	case roll < 45:
		return playMonsterRoom(c, InitGolem(), roomNumber)
	case roll < 54:
		return playMonsterRoom(c, InitTroll(), roomNumber)
	case roll < 61:
		return playMonsterRoom(c, InitSkeleton(), roomNumber)
	case roll < 66:
		return playMonsterRoom(c, InitWolf(), roomNumber)
	case roll < 70:
		return playMonsterRoom(c, InitWizard(), roomNumber)
	case roll < 78:
		if roomNumber >= 5 && roll < 75 {
			monster := scaleSpecialMonster(InitDragon(), roomNumber)
			return playMonsterRoom(c, monster, roomNumber)
		}
		if roomNumber >= 3 {
			monster := scaleSpecialMonster(InitDuck(), roomNumber)
			return playMonsterRoom(c, monster, roomNumber)
		}
		return playNpcRoom(c)
	case roll < 86:
		return playNpcRoom(c)
	case roll < 93:
		return playStatueRoom(c)
	default:
		return playChestRoom(c, random)
	}
}

func scaleSpecialMonster(monster Monster, roomNumber int) Monster {
	level := roomNumber/2 + 1
	monster.Level = level
	monster.MaxHP += (level - 1) * 25
	monster.CurrentHP = monster.MaxHP
	monster.Attack += (level - 1) * 3
	monster.Initiative += (level - 1) * 2
	monster.XPReward += (level - 1) * 40
	monster.GoldReward += (level - 1) * 20
	return monster
}

func playStatueRoom(c *Character) bool {
	fmt.Println("Vous trouvez une statue antique au centre de la salle.")
	fmt.Println("La statue émet une lumière sacrée et rétablit votre force.")
	c.CurrentHP = c.MaxHP
	c.Mana = c.MaxMana
	fmt.Printf("Vous récupérez tous vos PV et votre mana ! (%d / %d PV, %d / %d mana)\n", c.CurrentHP, c.MaxHP, c.Mana, c.MaxMana)
	WaitForEnter()
	return true
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

func playGhostRoom(c *Character, roomNumber int) bool {
	monster := InitGhostLevel("Fantôme", roomNumber)
	fmt.Printf("Un %s de niveau %d apparaît !\n", monster.Name, monster.Level)
	return normalFight(c, &monster, roomNumber-1)
}

func playMonsterRoom(c *Character, monster Monster, roomNumber int) bool {
	fmt.Printf("Vous rencontrez %s !\n", monster.Name)
	return normalFight(c, &monster, roomNumber-1)
}

func playBossRoom(c *Character, roomNumber int) bool {
	floor := dungeonFloor(roomNumber)
	monster := InitGoblinLevel("Boss gobelin", roomNumber+floor+1)
	monster.Name = "Boss gobelin"
	monster.Pattern = "goblin"
	monster.MaxHP += 60 + floor*25
	monster.CurrentHP = monster.MaxHP
	monster.Attack += 8 + floor*3
	monster.Initiative += 15 + floor*2
	monster.XPReward = c.MaxXP - c.CurrentXP + 50
	monster.GoldReward = 120 + floor*50
	fmt.Printf("BOSS DE L'ÉTAGE %d : un boss de niveau %d apparaît dans la salle !\n", floor, monster.Level)
	c.LastClearedRoom = roomNumber
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
			"Casque de gobelin", "Lame de gobelin",
			"Carapace de slime", "Bave de slime",
			"Capuche spectrale", "Lame spectrale",
			"Casque de golem", "Marteau de golem",
			"Peau de troll renforcée", "Massue de troll",
			"Plumes du canard", "Bec du canard",
			"Heaume squelette", "Épée squelette",
			"Fourrure du loup", "Crocs du loup",
			"Chapeau du sorcier", "Bâton maudit",
			"Écailles de dragon", "Griffe du dragon",
			"Tunique de gobelin", "Bottes de gobelin",
			"Tunique de slime", "Bottes de slime",
			"Tunique spectrale", "Bottes spectrales",
			"Tunique de golem", "Bottes de golem",
			"Tunique de troll", "Bottes de troll",
			"Tunique du canard", "Bottes du canard",
			"Tunique du squelette", "Bottes du squelette",
			"Tunique du loup", "Bottes du loup",
			"Tunique du sorcier", "Bottes du sorcier",
			"Tunique du dragon", "Bottes du dragon",
		}
		foundEquipment := equipment[random.Intn(len(equipment))]
		rarity := rollEquipmentRarity(c.LastClearedRoom + 1)
		foundEquipmentWithRarity := equipmentDisplayName(foundEquipment, rarity)
		fmt.Printf("Le coffre contient également : %s.\n", foundEquipmentWithRarity)
		c.AddOrMergeItem(foundEquipmentWithRarity, 1)
	}

	if random.Intn(2) == 0 {
		spells := []string{"Livre de Sort : Boule de Feu", "Livre de Sort : Soin", "Livre de Sort : Régénération", "Livre de Sort : Poison", "Livre de Sort : Brûlure", "Livre de Sort : Éclair", "Livre de Sort : Barrière Sacrée"}
		foundSpell := spells[random.Intn(len(spells))]
		fmt.Printf("Le coffre contient également : %s.\n", foundSpell)
		c.AddOrMergeItem(foundSpell, 1)
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
