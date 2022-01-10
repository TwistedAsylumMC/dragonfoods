package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type YellowCupcake struct{}

// AlwaysConsumable ...
func (YellowCupcake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (YellowCupcake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (YellowCupcake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (YellowCupcake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (YellowCupcake) Edible() bool {
	return true
}

// EncodeItem ...
func (YellowCupcake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:yellow_cupcake", 0
}

// Name ...
func (YellowCupcake) Name() string {
	return "YellowCupcake"
}

// Texture ...
func (YellowCupcake) Texture() image.Image {
	return textureFromName("yellowcupcake")
}
