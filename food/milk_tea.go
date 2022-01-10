package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MilkTea struct{}

// AlwaysConsumable ...
func (MilkTea) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MilkTea) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MilkTea) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (MilkTea) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MilkTea) Edible() bool {
	return true
}

// EncodeItem ...
func (MilkTea) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:milk_tea", 0
}

// Name ...
func (MilkTea) Name() string {
	return "Milk Tea"
}

// Texture ...
func (MilkTea) Texture() image.Image {
	return textureFromName("milktea3")
}
