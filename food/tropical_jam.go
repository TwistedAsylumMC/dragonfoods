package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type TropicalJam struct{}

// AlwaysConsumable ...
func (TropicalJam) AlwaysConsumable() bool {
	return false
}

// Category ...
func (TropicalJam) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (TropicalJam) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (TropicalJam) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (TropicalJam) Edible() bool {
	return true
}

// EncodeItem ...
func (TropicalJam) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:tropical_jam", 0
}

// Name ...
func (TropicalJam) Name() string {
	return "Tropical Jam"
}

// Texture ...
func (TropicalJam) Texture() image.Image {
	return textureFromName("61_jam")
}
