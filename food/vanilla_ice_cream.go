package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type VanillaIceCream struct{}

// AlwaysConsumable ...
func (VanillaIceCream) AlwaysConsumable() bool {
	return false
}

// Category ...
func (VanillaIceCream) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (VanillaIceCream) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (VanillaIceCream) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (VanillaIceCream) Edible() bool {
	return true
}

// EncodeItem ...
func (VanillaIceCream) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:vanilla_ice_cream", 0
}

// Name ...
func (VanillaIceCream) Name() string {
	return "Vanilla Ice Cream"
}

// Texture ...
func (VanillaIceCream) Texture() image.Image {
	return textureFromName("vanicecream")
}
