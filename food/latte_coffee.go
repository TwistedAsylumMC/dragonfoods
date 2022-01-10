package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type LatteCoffee struct{}

// AlwaysConsumable ...
func (LatteCoffee) AlwaysConsumable() bool {
	return false
}

// Category ...
func (LatteCoffee) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (LatteCoffee) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (LatteCoffee) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (LatteCoffee) Edible() bool {
	return true
}

// EncodeItem ...
func (LatteCoffee) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:latte_coffee", 0
}

// Name ...
func (LatteCoffee) Name() string {
	return "LatteCoffee"
}

// Texture ...
func (LatteCoffee) Texture() image.Image {
	return textureFromName("latte")
}
