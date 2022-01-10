package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Pecan struct{}

// AlwaysConsumable ...
func (Pecan) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Pecan) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Pecan) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Pecan) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Pecan) Edible() bool {
	return true
}

// EncodeItem ...
func (Pecan) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pecan", 0
}

// Name ...
func (Pecan) Name() string {
	return "Pecan"
}

// Texture ...
func (Pecan) Texture() image.Image {
	return textureFromName("pecan")
}
