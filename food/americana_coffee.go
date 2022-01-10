package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type AmericanaCoffee struct{}

// AlwaysConsumable ...
func (AmericanaCoffee) AlwaysConsumable() bool {
	return false
}

// Category ...
func (AmericanaCoffee) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (AmericanaCoffee) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (AmericanaCoffee) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (AmericanaCoffee) Edible() bool {
	return true
}

// EncodeItem ...
func (AmericanaCoffee) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:americana_coffee", 0
}

// Name ...
func (AmericanaCoffee) Name() string {
	return "AmericanaCoffee"
}

// Texture ...
func (AmericanaCoffee) Texture() image.Image {
	return textureFromName("starbucks_coffee")
}
