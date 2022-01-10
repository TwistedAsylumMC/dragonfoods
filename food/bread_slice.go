package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BreadSlice struct{}

// AlwaysConsumable ...
func (BreadSlice) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BreadSlice) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BreadSlice) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (BreadSlice) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BreadSlice) Edible() bool {
	return true
}

// EncodeItem ...
func (BreadSlice) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bread_slice", 0
}

// Name ...
func (BreadSlice) Name() string {
	return "Bread Slice"
}

// Texture ...
func (BreadSlice) Texture() image.Image {
	return textureFromName("loaf3")
}
