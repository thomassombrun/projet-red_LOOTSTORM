package library

import (
	"os"
	"strings"
	"testing"
)

func TestEquipmentRarityParsing(t *testing.T) {
	item := "Lame de gobelin [Épique]"
	if slot := slotForEquipment(item); slot != "Weapon" {
		t.Fatalf("slot attendu Weapon, obtenu %q", slot)
	}
	if rarity := parseItemRarity(item); rarity != RarityEpic {
		t.Fatalf("rarity attendue %q, obtenue %q", RarityEpic, rarity)
	}
	if base := normalizeItemName(item); base != "Lame de gobelin" {
		t.Fatalf("nom normalisé attendu %q, obtenu %q", "Lame de gobelin", base)
	}
}

func TestGenerateMonsterEquipment(t *testing.T) {
	drops := generateMonsterEquipment("goblin", 1)
	if len(drops) == 0 {
		t.Fatal("aucun équipement généré pour le gobelin")
	}
	valid := false
	for _, drop := range drops {
		if normalizeItemName(drop) == "Casque de gobelin" || normalizeItemName(drop) == "Lame de gobelin" || normalizeItemName(drop) == "Tunique de gobelin" || normalizeItemName(drop) == "Bottes de gobelin" {
			valid = true
		}
	}
	if !valid {
		t.Fatal("le gobelin doit pouvoir dropper un équipement de son propre pool")
	}
}

func TestRarityRollIsNotEmpty(t *testing.T) {
	if rarity := rollEquipmentRarity(8); rarity == "" {
		t.Fatal("une rareté doit être retournée")
	}
}

func TestEquipmentKeepsRarityOnEquip(t *testing.T) {
	c := InitCharacter("Test", "Guerrier", 100)
	c.Inventory = []Item{{Name: "Heaume squelette", Quantity: 1, Rarity: RarityEpic}}
	EquipItem(&c, itemDisplayName(c.Inventory[0]), 0)
	if !strings.Contains(c.Equip.Helmet, "[Épique]") {
		t.Fatalf("la rareté doit être conservée après l'équipement, obtenu %q", c.Equip.Helmet)
	}
	if equipmentHPBonus("Heaume squelette") == 0 {
		t.Fatal("un heaume squelette doit donner des PV")
	}
}

func TestMarteauDeGolemIsWeapon(t *testing.T) {
	if slotForEquipment("Marteau de golem [Épique]") != "Weapon" {
		t.Fatal("le marteau de golem doit être classé en Weapon")
	}
	if weaponDamageBonus("Marteau de golem") <= 0 {
		t.Fatal("le marteau de golem doit donner de l'attaque")
	}
	if equipmentHPBonus("Marteau de golem") != 0 {
		t.Fatal("le marteau de golem ne doit pas donner de PV")
	}
}

func TestWeaponBonusCountsInAttackDamage(t *testing.T) {
	c := InitCharacter("Test", "Guerrier", 100)
	withoutWeapon := BasicAttackDamage(&c)
	c.Equip.Weapon = "Marteau de golem [Rare]"
	c.Equip.WeaponDamage = 18
	withWeapon := BasicAttackDamage(&c)
	if withWeapon-withoutWeapon != 18 {
		t.Fatalf("le bonus d'arme devrait ajouter 18 dégâts, différence obtenue : %d", withWeapon-withoutWeapon)
	}
}

func TestReplaceEquipmentSameCategory(t *testing.T) {
	c := InitCharacter("Test", "Guerrier", 100)
	c.Equip.Boots = "Bottes de slime [Commun]"
	c.Inventory = []Item{{Name: "Bottes spectrales", Quantity: 1, Rarity: RarityRare}}
	oldLen := len(c.Inventory)
	oldEquipped := c.Equip.Boots
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	w.WriteString("1\n")
	w.Close()
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }()
	EquipItem(&c, itemDisplayName(c.Inventory[0]), 0)
	if c.Equip.Boots == oldEquipped {
		t.Fatal("l'ancien équipement devait être remplacé par le nouveau")
	}
	if len(c.Inventory) != oldLen {
		t.Fatal("l'ancien équipement n'a pas bien été remis dans l'inventaire")
	}
	if !strings.Contains(c.Equip.Boots, "Bottes spectrales") {
		t.Fatalf("la nouvelle botte n'est pas équipée : %q", c.Equip.Boots)
	}
}
