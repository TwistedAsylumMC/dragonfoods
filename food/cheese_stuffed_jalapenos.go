package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CheeseStuffedJalapenos struct{}

// AlwaysConsumable ...
func (CheeseStuffedJalapenos) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CheeseStuffedJalapenos) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CheeseStuffedJalapenos) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (CheeseStuffedJalapenos) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CheeseStuffedJalapenos) Edible() bool {
	return true
}

// EncodeItem ...
func (CheeseStuffedJalapenos) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cheese_stuffed_jalapenos", 0
}

// Name ...
func (CheeseStuffedJalapenos) Name() string {
	return "Cheese Stuffed Jalapenos"
}

// Texture ...
func (CheeseStuffedJalapenos) Texture() image.Image {
	return textureFromName("stujalapenos")
}
