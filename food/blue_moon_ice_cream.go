package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueMoonIceCream struct{}

// AlwaysConsumable ...
func (BlueMoonIceCream) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueMoonIceCream) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueMoonIceCream) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueMoonIceCream) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueMoonIceCream) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueMoonIceCream) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_moon_ice_cream", 0
}

// Name ...
func (BlueMoonIceCream) Name() string {
	return "Blue Moon Ice Cream"
}

// Texture ...
func (BlueMoonIceCream) Texture() image.Image {
	return textureFromName("iccream3")
}
