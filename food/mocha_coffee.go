package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MochaCoffee struct{}

// AlwaysConsumable ...
func (MochaCoffee) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MochaCoffee) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MochaCoffee) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (MochaCoffee) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MochaCoffee) Edible() bool {
	return true
}

// EncodeItem ...
func (MochaCoffee) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:mocha_coffee", 0
}

// Name ...
func (MochaCoffee) Name() string {
	return "MochaCoffee"
}

// Texture ...
func (MochaCoffee) Texture() image.Image {
	return textureFromName("mocha_coffee")
}
