package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type YellowBellPepper struct{}

// AlwaysConsumable ...
func (YellowBellPepper) AlwaysConsumable() bool {
	return false
}

// Category ...
func (YellowBellPepper) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (YellowBellPepper) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (YellowBellPepper) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (YellowBellPepper) Edible() bool {
	return true
}

// EncodeItem ...
func (YellowBellPepper) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:yellow_bell_pepper", 0
}

// Name ...
func (YellowBellPepper) Name() string {
	return "YellowBellPepper"
}

// Texture ...
func (YellowBellPepper) Texture() image.Image {
	return textureFromName("pepper_bell_02")
}
