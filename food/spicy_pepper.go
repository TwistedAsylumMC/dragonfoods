package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SpicyPepper struct{}

// AlwaysConsumable ...
func (SpicyPepper) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SpicyPepper) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SpicyPepper) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (SpicyPepper) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SpicyPepper) Edible() bool {
	return true
}

// EncodeItem ...
func (SpicyPepper) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:spicy_pepper", 0
}

// Name ...
func (SpicyPepper) Name() string {
	return "SpicyPepper"
}

// Texture ...
func (SpicyPepper) Texture() image.Image {
	return textureFromName("pepper2")
}
