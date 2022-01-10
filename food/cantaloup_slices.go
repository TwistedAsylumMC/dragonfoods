package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CantaloupSlices struct{}

// AlwaysConsumable ...
func (CantaloupSlices) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CantaloupSlices) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CantaloupSlices) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (CantaloupSlices) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CantaloupSlices) Edible() bool {
	return true
}

// EncodeItem ...
func (CantaloupSlices) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cantaloup_slices", 0
}

// Name ...
func (CantaloupSlices) Name() string {
	return "CantaloupSlices"
}

// Texture ...
func (CantaloupSlices) Texture() image.Image {
	return textureFromName("meloncantaloupe")
}
