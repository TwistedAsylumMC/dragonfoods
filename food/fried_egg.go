package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type FriedEgg struct{}

// AlwaysConsumable ...
func (FriedEgg) AlwaysConsumable() bool {
	return false
}

// Category ...
func (FriedEgg) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (FriedEgg) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (FriedEgg) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (FriedEgg) Edible() bool {
	return true
}

// EncodeItem ...
func (FriedEgg) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:fried_egg", 0
}

// Name ...
func (FriedEgg) Name() string {
	return "Fried Egg"
}

// Texture ...
func (FriedEgg) Texture() image.Image {
	return textureFromName("friedegg")
}
