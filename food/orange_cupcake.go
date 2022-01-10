package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type OrangeCupcake struct{}

// AlwaysConsumable ...
func (OrangeCupcake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (OrangeCupcake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (OrangeCupcake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (OrangeCupcake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (OrangeCupcake) Edible() bool {
	return true
}

// EncodeItem ...
func (OrangeCupcake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:orange_cupcake", 0
}

// Name ...
func (OrangeCupcake) Name() string {
	return "Orange Cupcake"
}

// Texture ...
func (OrangeCupcake) Texture() image.Image {
	return textureFromName("orangecupcake")
}
