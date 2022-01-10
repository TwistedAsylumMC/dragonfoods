package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PinkCake struct{}

// AlwaysConsumable ...
func (PinkCake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PinkCake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PinkCake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (PinkCake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PinkCake) Edible() bool {
	return true
}

// EncodeItem ...
func (PinkCake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pink_cake", 0
}

// Name ...
func (PinkCake) Name() string {
	return "PinkCake"
}

// Texture ...
func (PinkCake) Texture() image.Image {
	return textureFromName("bakedpinkcake")
}
