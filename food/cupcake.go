package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Cupcake struct{}

// AlwaysConsumable ...
func (Cupcake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Cupcake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Cupcake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Cupcake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Cupcake) Edible() bool {
	return true
}

// EncodeItem ...
func (Cupcake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cupcake", 0
}

// Name ...
func (Cupcake) Name() string {
	return "Cupcake"
}

// Texture ...
func (Cupcake) Texture() image.Image {
	return textureFromName("cupcake")
}
