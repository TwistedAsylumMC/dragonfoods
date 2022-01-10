package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Kumquat struct{}

// AlwaysConsumable ...
func (Kumquat) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Kumquat) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Kumquat) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Kumquat) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Kumquat) Edible() bool {
	return true
}

// EncodeItem ...
func (Kumquat) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:kumquat", 0
}

// Name ...
func (Kumquat) Name() string {
	return "Kumquat"
}

// Texture ...
func (Kumquat) Texture() image.Image {
	return textureFromName("kumquat_01")
}
