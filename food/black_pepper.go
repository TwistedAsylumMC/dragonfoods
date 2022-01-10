package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlackPepper struct{}

// AlwaysConsumable ...
func (BlackPepper) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlackPepper) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlackPepper) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlackPepper) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlackPepper) Edible() bool {
	return true
}

// EncodeItem ...
func (BlackPepper) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:black_pepper", 0
}

// Name ...
func (BlackPepper) Name() string {
	return "BlackPepper"
}

// Texture ...
func (BlackPepper) Texture() image.Image {
	return textureFromName("pepper")
}
