package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Peanut struct{}

// AlwaysConsumable ...
func (Peanut) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Peanut) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Peanut) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Peanut) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Peanut) Edible() bool {
	return true
}

// EncodeItem ...
func (Peanut) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:peanut", 0
}

// Name ...
func (Peanut) Name() string {
	return "Peanut"
}

// Texture ...
func (Peanut) Texture() image.Image {
	return textureFromName("peanut")
}
