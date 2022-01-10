package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PineappleSlice struct{}

// AlwaysConsumable ...
func (PineappleSlice) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PineappleSlice) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PineappleSlice) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (PineappleSlice) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PineappleSlice) Edible() bool {
	return true
}

// EncodeItem ...
func (PineappleSlice) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pineapple_slice", 0
}

// Name ...
func (PineappleSlice) Name() string {
	return "PineappleSlice"
}

// Texture ...
func (PineappleSlice) Texture() image.Image {
	return textureFromName("pineappleslice")
}
