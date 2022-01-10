package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChoclateBar struct{}

// AlwaysConsumable ...
func (ChoclateBar) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChoclateBar) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChoclateBar) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChoclateBar) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChoclateBar) Edible() bool {
	return true
}

// EncodeItem ...
func (ChoclateBar) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:choclate_bar", 0
}

// Name ...
func (ChoclateBar) Name() string {
	return "ChoclateBar"
}

// Texture ...
func (ChoclateBar) Texture() image.Image {
	return textureFromName("milk_chocolate")
}
