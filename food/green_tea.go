package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GreenTea struct{}

// AlwaysConsumable ...
func (GreenTea) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GreenTea) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GreenTea) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (GreenTea) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GreenTea) Edible() bool {
	return true
}

// EncodeItem ...
func (GreenTea) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:green_tea", 0
}

// Name ...
func (GreenTea) Name() string {
	return "Green Tea"
}

// Texture ...
func (GreenTea) Texture() image.Image {
	return textureFromName("tea_03")
}
