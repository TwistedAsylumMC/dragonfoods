package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type HighClassSteak struct{}

// AlwaysConsumable ...
func (HighClassSteak) AlwaysConsumable() bool {
	return false
}

// Category ...
func (HighClassSteak) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (HighClassSteak) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (HighClassSteak) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (HighClassSteak) Edible() bool {
	return true
}

// EncodeItem ...
func (HighClassSteak) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:high_class_steak", 0
}

// Name ...
func (HighClassSteak) Name() string {
	return "High Class Steak"
}

// Texture ...
func (HighClassSteak) Texture() image.Image {
	return textureFromName("steak2")
}
