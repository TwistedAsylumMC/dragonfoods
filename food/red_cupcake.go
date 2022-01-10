package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RedCupcake struct{}

// AlwaysConsumable ...
func (RedCupcake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RedCupcake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RedCupcake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (RedCupcake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RedCupcake) Edible() bool {
	return true
}

// EncodeItem ...
func (RedCupcake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:red_cupcake", 0
}

// Name ...
func (RedCupcake) Name() string {
	return "Red Cupcake"
}

// Texture ...
func (RedCupcake) Texture() image.Image {
	return textureFromName("recupcake")
}
