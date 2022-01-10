package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type FruitBowel struct{}

// AlwaysConsumable ...
func (FruitBowel) AlwaysConsumable() bool {
	return false
}

// Category ...
func (FruitBowel) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (FruitBowel) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (FruitBowel) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (FruitBowel) Edible() bool {
	return true
}

// EncodeItem ...
func (FruitBowel) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:fruit_bowel", 0
}

// Name ...
func (FruitBowel) Name() string {
	return "Fruit Bowel"
}

// Texture ...
func (FruitBowel) Texture() image.Image {
	return textureFromName("fruitbowel")
}
