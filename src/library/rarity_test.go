package library

import "testing"

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
