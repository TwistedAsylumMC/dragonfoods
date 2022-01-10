package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Pickle struct{}

// AlwaysConsumable ...
func (Pickle) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Pickle) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Pickle) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Pickle) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Pickle) Edible() bool {
	return true
}

// EncodeItem ...
func (Pickle) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pickle", 0
}

// Name ...
func (Pickle) Name() string {
	return "Pickle"
}

// Texture ...
func (Pickle) Texture() image.Image {
	return textureFromName("pickle")
}
