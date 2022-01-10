package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Sub struct{}

// AlwaysConsumable ...
func (Sub) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Sub) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Sub) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Sub) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Sub) Edible() bool {
	return true
}

// EncodeItem ...
func (Sub) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sub", 0
}

// Name ...
func (Sub) Name() string {
	return "Sub"
}

// Texture ...
func (Sub) Texture() image.Image {
	return textureFromName("sub")
}
