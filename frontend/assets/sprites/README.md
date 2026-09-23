# Sprites Lootstorm

Place tes images PNG dans ce dossier avec les noms suivants.

## Classes des personnages créés

- `class_guerrier.png`
- `class_mage.png`
- `class_archer.png`
- `class_assassin.png`
- `class_chevalier.png`
- `class_samourai.png`
- `class_clerc.png`
- `class_barbare.png`
- `class_invocateur.png`

## Heros predefinis

- `hero_himiko_toga.png`
- `hero_link.png`
- `hero_frieren.png`
- `hero_musashi.png`
- `hero_elizabeth.png`
- `hero_guts.png`
- `hero_sung_jin_woo.png`
- `hero_colley.png`

## Ennemis

- `enemy_goblin.png`
- `enemy_slime.png`
- `enemy_ghost.png`
- `enemy_golem.png`
- `enemy_troll.png`
- `enemy_skeleton.png`
- `enemy_wolf.png`
- `enemy_wizard.png`
- `enemy_duck.png`
- `enemy_dragon.png`
- `enemy_summon.png`

Les fichiers doivent etre au format PNG. La transparence est recommandee. Les sprites peuvent avoir n'importe quelle taille : ils seront redimensionnes automatiquement dans une zone de 128x128 pixels.

Lance le jeu depuis la racine du projet avec :

```bash
go run ./frontend
```

Les sprites absents utilisent d'abord le sprite de classe, puis le dessin de secours integre au jeu. Cela permet de ne fournir qu'une partie des images.
