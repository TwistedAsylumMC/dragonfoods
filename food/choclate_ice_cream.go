package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChoclateIceCream struct{}

// AlwaysConsumable ...
func (ChoclateIceCream) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChoclateIceCream) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChoclateIceCream) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChoclateIceCream) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChoclateIceCream) Edible() bool {
	return true
}

// EncodeItem ...
func (ChoclateIceCream) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:choclate_ice_cream", 0
}

// Name ...
func (ChoclateIceCream) Name() string {
	return "ChoclateIceCream"
}

// Texture ...
func (ChoclateIceCream) Texture() image.Image {
	return textureFromName("choclate_ice-cream")
}
