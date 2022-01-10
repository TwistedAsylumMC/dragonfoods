package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Peach struct{}

// AlwaysConsumable ...
func (Peach) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Peach) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Peach) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Peach) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Peach) Edible() bool {
	return true
}

// EncodeItem ...
func (Peach) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:peach", 0
}

// Name ...
func (Peach) Name() string {
	return "Peach"
}

// Texture ...
func (Peach) Texture() image.Image {
	return textureFromName("peach2")
}
