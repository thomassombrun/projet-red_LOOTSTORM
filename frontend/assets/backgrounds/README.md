# Fond de jeu

Ajoute ton image ici sous le nom exact :

```text
dungeon.png
```

## Fond special de combat

Pour afficher une image uniquement pendant l'aventure et les combats d'entrainement, ajoute aussi :

```text
combat.png
```

Ce fond est utilise pendant le combat principal, le choix des sorts et le choix des objets. Il ne s'affiche pas dans les menus, la creation de personnage ou l'inventaire.

## Fonds des services

Tu peux aussi ajouter ces images pour personnaliser chaque service :

```text
merchant.png
forge.png
enchanter.png
```

- `merchant.png` : écran du marchand, achats et ventes ;
- `forge.png` : écran du forgeron ;
- `enchanter.png` : écran de l'enchanteur.

Chaque image est utilisée uniquement dans son écran. Si elle est absente, `dungeon.png` est utilisé automatiquement.

Recommandations :

- PNG ou JPG converti en PNG ;
- resolution ideale : 960x514 pour remplir la zone sous le header, ou 1920x1028 ;
- ratio proche de 16:8.57 pour eviter le recadrage ;
- sans texte integre dans l'image ;
- zone centrale et basse peu contrastee pour garder les menus lisibles ;
- elements importants places sur les bords.

Le jeu redimensionne automatiquement les images et ajoute un voile sombre leger. Si `dungeon.png` ou `combat.png` est absent, le fond procedural est utilise automatiquement.

Lance le jeu depuis la racine :

```bash
go run ./frontend
```
