package main
package main

import (
	"testing"
)

func TestEmbeddedSprites(t *testing.T) {
	for _, name := range []string{"Himiko Toga", "Link", "Frieren"} {
		path := "assets/characters/" + slugify(name) + ".png"
		if _, err := spriteAssets.ReadFile(path); err != nil {
			t.Errorf("%s -> %s: %v", name, path, err)
		} else {
			t.Logf("%s -> %s OK", name, path)
		}
	}
}
