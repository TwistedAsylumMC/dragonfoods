package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GreenCupcake struct{}

// AlwaysConsumable ...
func (GreenCupcake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GreenCupcake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GreenCupcake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (GreenCupcake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GreenCupcake) Edible() bool {
	return true
}

// EncodeItem ...
func (GreenCupcake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:green_cupcake", 0
}

// Name ...
func (GreenCupcake) Name() string {
	return "Green Cupcake"
}

// Texture ...
func (GreenCupcake) Texture() image.Image {
	return textureFromName("greencupcake")
}
