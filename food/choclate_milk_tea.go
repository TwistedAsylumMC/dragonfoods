package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChoclateMilkTea struct{}

// AlwaysConsumable ...
func (ChoclateMilkTea) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChoclateMilkTea) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChoclateMilkTea) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChoclateMilkTea) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChoclateMilkTea) Edible() bool {
	return true
}

// EncodeItem ...
func (ChoclateMilkTea) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:choclate_milk_tea", 0
}

// Name ...
func (ChoclateMilkTea) Name() string {
	return "Choclate Milk Tea"
}

// Texture ...
func (ChoclateMilkTea) Texture() image.Image {
	return textureFromName("milktea")
}
