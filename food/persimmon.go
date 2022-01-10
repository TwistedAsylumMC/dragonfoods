package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Persimmon struct{}

// AlwaysConsumable ...
func (Persimmon) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Persimmon) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Persimmon) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Persimmon) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Persimmon) Edible() bool {
	return true
}

// EncodeItem ...
func (Persimmon) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:persimmon", 0
}

// Name ...
func (Persimmon) Name() string {
	return "Persimmon"
}

// Texture ...
func (Persimmon) Texture() image.Image {
	return textureFromName("persimmon")
}
