package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MilkChocateBar struct{}

// AlwaysConsumable ...
func (MilkChocateBar) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MilkChocateBar) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MilkChocateBar) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (MilkChocateBar) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MilkChocateBar) Edible() bool {
	return true
}

// EncodeItem ...
func (MilkChocateBar) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:milk_chocate_bar", 0
}

// Name ...
func (MilkChocateBar) Name() string {
	return "MilkChocateBar"
}

// Texture ...
func (MilkChocateBar) Texture() image.Image {
	return textureFromName("milk_choclate2")
}
