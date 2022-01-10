package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChoclateChips struct{}

// AlwaysConsumable ...
func (ChoclateChips) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChoclateChips) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChoclateChips) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChoclateChips) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (ChoclateChips) Edible() bool {
	return true
}

// EncodeItem ...
func (ChoclateChips) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:choclate_chips", 0
}

// Name ...
func (ChoclateChips) Name() string {
	return "Choclate Chips"
}

// Texture ...
func (ChoclateChips) Texture() image.Image {
	return textureFromName("choclatechips")
}
