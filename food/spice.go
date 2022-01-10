package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Spice struct{}

// AlwaysConsumable ...
func (Spice) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Spice) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Spice) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Spice) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Spice) Edible() bool {
	return true
}

// EncodeItem ...
func (Spice) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:spice", 0
}

// Name ...
func (Spice) Name() string {
	return "Spice"
}

// Texture ...
func (Spice) Texture() image.Image {
	return textureFromName("spice")
}
