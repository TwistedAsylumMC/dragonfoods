package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Coconut struct{}

// AlwaysConsumable ...
func (Coconut) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Coconut) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Coconut) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Coconut) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Coconut) Edible() bool {
	return true
}

// EncodeItem ...
func (Coconut) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:coconut", 0
}

// Name ...
func (Coconut) Name() string {
	return "Coconut"
}

// Texture ...
func (Coconut) Texture() image.Image {
	return textureFromName("coconut_02")
}
