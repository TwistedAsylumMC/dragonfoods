package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Kale struct{}

// AlwaysConsumable ...
func (Kale) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Kale) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Kale) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Kale) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Kale) Edible() bool {
	return true
}

// EncodeItem ...
func (Kale) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:kale", 0
}

// Name ...
func (Kale) Name() string {
	return "Kale"
}

// Texture ...
func (Kale) Texture() image.Image {
	return textureFromName("kale_01")
}
