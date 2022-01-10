package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Peanutbuttercookie struct{}

// AlwaysConsumable ...
func (Peanutbuttercookie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Peanutbuttercookie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Peanutbuttercookie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Peanutbuttercookie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Peanutbuttercookie) Edible() bool {
	return true
}

// EncodeItem ...
func (Peanutbuttercookie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:peanutbuttercookie", 0
}

// Name ...
func (Peanutbuttercookie) Name() string {
	return "Peanutbuttercookie"
}

// Texture ...
func (Peanutbuttercookie) Texture() image.Image {
	return textureFromName("peanutbuttercookie")
}
