package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type IrishCoffee struct{}

// AlwaysConsumable ...
func (IrishCoffee) AlwaysConsumable() bool {
	return false
}

// Category ...
func (IrishCoffee) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (IrishCoffee) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (IrishCoffee) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (IrishCoffee) Edible() bool {
	return true
}

// EncodeItem ...
func (IrishCoffee) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:irish_coffee", 0
}

// Name ...
func (IrishCoffee) Name() string {
	return "IrishCoffee"
}

// Texture ...
func (IrishCoffee) Texture() image.Image {
	return textureFromName("irish_coffee")
}
