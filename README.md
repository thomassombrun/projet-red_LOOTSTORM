# projet-red_LOOTSTORM

Ceci est un jeu en langage GO, nous nous sommes inspirés des jeux dans le style Donjons & Dragons pour le thème et RPG pour le gameplay (style Roguelite)

## Lancer juste le backend

Depuis le terminal :

```bash
go run ./src
```

Utilisez les numéros de 1 à 9 pour faire vos choix et 0 pour revenir en arrière

## Lancer le frontend Ebiten

Depuis la racine du projet :

```bash
go run ./frontend
```

Utilisez les fleches pour naviguer, `Entree` pour confirmer et `Echap` pour revenir ou quitter.

Le frontend affiche actuellement la selection des heros, le menu principal et la fiche du personnage. Les systemes d'aventure, de combat, d'inventaire, de marchand, de forge et d'enchantement restent dans l'interface console existante et seront branches progressivement sur des ecrans Ebiten.