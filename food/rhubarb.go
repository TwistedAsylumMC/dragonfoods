package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Rhubarb struct{}

// AlwaysConsumable ...
func (Rhubarb) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Rhubarb) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Rhubarb) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Rhubarb) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Rhubarb) Edible() bool {
	return true
}

// EncodeItem ...
func (Rhubarb) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:rhubarb", 0
}

// Name ...
func (Rhubarb) Name() string {
	return "Rhubarb"
}

// Texture ...
func (Rhubarb) Texture() image.Image {
	return textureFromName("rhubarb_01")
}
