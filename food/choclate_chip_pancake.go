package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChoclateChipPancake struct{}

// AlwaysConsumable ...
func (ChoclateChipPancake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChoclateChipPancake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChoclateChipPancake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChoclateChipPancake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChoclateChipPancake) Edible() bool {
	return true
}

// EncodeItem ...
func (ChoclateChipPancake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:choclate_chip_pancake", 0
}

// Name ...
func (ChoclateChipPancake) Name() string {
	return "ChoclateChipPancake"
}

// Texture ...
func (ChoclateChipPancake) Texture() image.Image {
	return textureFromName("pancake3")
}
