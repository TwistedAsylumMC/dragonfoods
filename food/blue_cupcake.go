package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueCupcake struct{}

// AlwaysConsumable ...
func (BlueCupcake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueCupcake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueCupcake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueCupcake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueCupcake) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueCupcake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_cupcake", 0
}

// Name ...
func (BlueCupcake) Name() string {
	return "Blue Cupcake"
}

// Texture ...
func (BlueCupcake) Texture() image.Image {
	return textureFromName("bluecupcake")
}
