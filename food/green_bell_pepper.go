package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GreenBellPepper struct{}

// AlwaysConsumable ...
func (GreenBellPepper) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GreenBellPepper) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GreenBellPepper) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (GreenBellPepper) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GreenBellPepper) Edible() bool {
	return true
}

// EncodeItem ...
func (GreenBellPepper) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:green_bell_pepper", 0
}

// Name ...
func (GreenBellPepper) Name() string {
	return "GreenBellPepper"
}

// Texture ...
func (GreenBellPepper) Texture() image.Image {
	return textureFromName("pepper_bell_01")
}
