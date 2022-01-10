package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type HardBoiledEgg struct{}

// AlwaysConsumable ...
func (HardBoiledEgg) AlwaysConsumable() bool {
	return false
}

// Category ...
func (HardBoiledEgg) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (HardBoiledEgg) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (HardBoiledEgg) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (HardBoiledEgg) Edible() bool {
	return true
}

// EncodeItem ...
func (HardBoiledEgg) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:hard_boiled_egg", 0
}

// Name ...
func (HardBoiledEgg) Name() string {
	return "Hard Boiled Egg"
}

// Texture ...
func (HardBoiledEgg) Texture() image.Image {
	return textureFromName("hard_boiled_egg")
}
