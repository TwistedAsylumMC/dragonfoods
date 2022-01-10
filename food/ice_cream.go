package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type IceCream struct{}

// AlwaysConsumable ...
func (IceCream) AlwaysConsumable() bool {
	return false
}

// Category ...
func (IceCream) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (IceCream) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (IceCream) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (IceCream) Edible() bool {
	return true
}

// EncodeItem ...
func (IceCream) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:ice_cream", 0
}

// Name ...
func (IceCream) Name() string {
	return "IceCream"
}

// Texture ...
func (IceCream) Texture() image.Image {
	return textureFromName("icecream")
}
