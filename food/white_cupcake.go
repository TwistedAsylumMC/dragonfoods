package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type WhiteCupcake struct{}

// AlwaysConsumable ...
func (WhiteCupcake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (WhiteCupcake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (WhiteCupcake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (WhiteCupcake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (WhiteCupcake) Edible() bool {
	return true
}

// EncodeItem ...
func (WhiteCupcake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:white_cupcake", 0
}

// Name ...
func (WhiteCupcake) Name() string {
	return "White Cupcake"
}

// Texture ...
func (WhiteCupcake) Texture() image.Image {
	return textureFromName("whitecupcake")
}
