package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Lasagna struct{}

// AlwaysConsumable ...
func (Lasagna) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Lasagna) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Lasagna) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Lasagna) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Lasagna) Edible() bool {
	return true
}

// EncodeItem ...
func (Lasagna) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:lasagna", 0
}

// Name ...
func (Lasagna) Name() string {
	return "Lasagna"
}

// Texture ...
func (Lasagna) Texture() image.Image {
	return textureFromName("lasagna")
}
