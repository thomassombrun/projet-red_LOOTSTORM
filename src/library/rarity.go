package library

import (
	"fmt"
	"math/rand"
	"strings"
)

type Rarity string

const (
	RarityCommon    Rarity = "Commun"
	RarityUncommon  Rarity = "Peu commun"
	RarityRare      Rarity = "Rare"
	RarityEpic      Rarity = "Épique"
	RarityLegendary Rarity = "Légendaire"
	RarityMythic    Rarity = "Mythique"
)

func (r Rarity) String() string {
	if r == "" {
		return string(RarityCommon)
	}
	return string(r)
}

func normalizeItemName(name string) string {
	trimmed := strings.TrimSpace(name)
	if idx := strings.LastIndex(trimmed, "["); idx >= 0 && strings.HasSuffix(trimmed, "]") {
		trimmed = strings.TrimSpace(trimmed[:idx])
	}
	return trimmed
}

func parseItemRarity(name string) Rarity {
	trimmed := strings.TrimSpace(name)
	if idx := strings.LastIndex(trimmed, "["); idx >= 0 && strings.HasSuffix(trimmed, "]") {
		value := strings.TrimSpace(trimmed[idx+1 : len(trimmed)-1])
		switch value {
		case "Commun", "commun":
			return RarityCommon
		case "Peu commun", "peu commun":
			return RarityUncommon
		case "Rare", "rare":
			return RarityRare
		case "Épique", "epique", "Epic", "epic":
			return RarityEpic
		case "Légendaire", "legendaire", "Legendary", "legendary":
			return RarityLegendary
		case "Mythique", "mythique", "Mythic", "mythic":
			return RarityMythic
		}
	}
	return RarityCommon
}

func itemDisplayName(item Item) string {
	if item.Rarity == "" {
		return item.Name
	}
	return fmt.Sprintf("%s [%s]", item.Name, item.Rarity)
}

func ItemDisplayName(item Item) string {
	return itemDisplayName(item)
}

func equipmentDisplayName(name string, rarity Rarity) string {
	base := normalizeItemName(name)
	if rarity == "" {
		return base
	}
	return fmt.Sprintf("%s [%s]", base, rarity)
}

func itemStatSummary(itemName string) string {
	base := normalizeItemName(itemName)
	var stats []string
	if hp := equipmentHPBonus(base); hp > 0 {
		stats = append(stats, fmt.Sprintf("+%d PV", hp))
	}
	if dmg := weaponDamageBonus(base); dmg > 0 {
		stats = append(stats, fmt.Sprintf("+%d ATQ", dmg))
	}
	if len(stats) == 0 {
		return ""
	}
	return " (" + strings.Join(stats, " / ") + ")"
}
func ItemStatSummary(itemName string) string {
	return itemStatSummary(itemName)
}

func rollEquipmentRarity(dungeonTier int) Rarity {
	if dungeonTier < 1 {
		dungeonTier = 1
	}

	weights := map[Rarity]int{
		RarityCommon:    50 - dungeonTier*2,
		RarityUncommon:  26 + dungeonTier,
		RarityRare:      15 + dungeonTier*2,
		RarityEpic:      6 + dungeonTier,
		RarityLegendary: 2 + dungeonTier/2,
		RarityMythic:    1 + dungeonTier/3,
	}

	if weights[RarityCommon] < 5 {
		weights[RarityCommon] = 5
	}

	roll := rand.Intn(100) + 1
	threshold := 0
	for _, rarity := range []Rarity{RarityCommon, RarityUncommon, RarityRare, RarityEpic, RarityLegendary, RarityMythic} {
		threshold += weights[rarity]
		if roll <= threshold {
			return rarity
		}
	}
	return RarityCommon
}

func rarityMultiplier(rarity Rarity) float64 {
	switch rarity {
	case RarityCommon:
		return 1.0
	case RarityUncommon:
		return 1.15
	case RarityRare:
		return 1.35
	case RarityEpic:
		return 1.6
	case RarityLegendary:
		return 1.9
	case RarityMythic:
		return 2.3
	default:
		return 1.0
	}
}

func monsterEquipmentPool(pattern string) []string {
	switch pattern {
	case "goblin":
		return []string{"Casque de gobelin", "Lame de gobelin", "Tunique de gobelin", "Bottes de gobelin"}
	case "slime":
		return []string{"Carapace de slime", "Bave de slime", "Tunique de slime", "Bottes de slime"}
	case "ghost":
		return []string{"Capuche spectrale", "Lame spectrale", "Tunique spectrale", "Bottes spectrales"}
	case "golem":
		return []string{"Casque de golem", "Marteau de golem", "Tunique de golem", "Bottes de golem"}
	case "troll":
		return []string{"Peau de troll renforcée", "Massue de troll", "Tunique de troll", "Bottes de troll"}
	case "duck":
		return []string{"Plumes du canard", "Bec du canard", "Tunique du canard", "Bottes du canard"}
	case "skeleton":
		return []string{"Heaume squelette", "Épée squelette", "Tunique du squelette", "Bottes du squelette"}
	case "wolf":
		return []string{"Fourrure du loup", "Crocs du loup", "Tunique du loup", "Bottes du loup"}
	case "wizard":
		return []string{"Chapeau du sorcier", "Bâton maudit", "Tunique du sorcier", "Bottes du sorcier"}
	case "dragon":
		return []string{"Écailles de dragon", "Griffe du dragon", "Tunique du dragon", "Bottes du dragon"}
	default:
		return nil
	}
}

func generateMonsterEquipment(pattern string, dungeonTier int) []string {
	pool := monsterEquipmentPool(pattern)
	if len(pool) == 0 {
		return nil
	}

	count := 1
	if rand.Intn(100) < 25 && dungeonTier > 1 {
		count = 2
	}
	if count > len(pool) {
		count = len(pool)
	}

	result := make([]string, 0, count)
	used := map[string]bool{}
	for len(result) < count {
		item := pool[rand.Intn(len(pool))]
		if used[item] {
			continue
		}
		used[item] = true
		result = append(result, item)
	}
	return result
}

func slotForEquipment(itemName string) string {
	normalized := normalizeItemName(itemName)
	switch normalized {
	case "Chapeau de l'aventurier", "Casque de gobelin", "Carapace de slime", "Capuche spectrale", "Casque de golem", "Peau de troll renforcée", "Plumes du canard", "Heaume squelette", "Fourrure du loup", "Chapeau du sorcier", "Écailles de dragon":
		return "Helmet"
	case "Tunique de l'aventurier", "Tunique de gobelin", "Tunique de slime", "Tunique spectrale", "Tunique de golem", "Tunique de troll", "Tunique du canard", "Tunique du squelette", "Tunique du loup", "Tunique du sorcier", "Tunique du dragon":
		return "Chestplate"
	case "Bottes de l'aventurier", "Bottes de gobelin", "Bottes de slime", "Bottes spectrales", "Bottes de golem", "Bottes de troll", "Bottes du canard", "Bottes du squelette", "Bottes du loup", "Bottes du sorcier", "Bottes du dragon":
		return "Boots"
	case "Dague de l'assassin", "Arc du chasseur", "Marteau du guerrier", "Épée du chevalier", "Lame de gobelin", "Bave de slime", "Lame spectrale", "Marteau de golem", "Massue de troll", "Bec du canard", "Épée squelette", "Crocs du loup", "Bâton maudit", "Griffe du dragon":
		return "Weapon"
	default:
		return ""
	}
}
