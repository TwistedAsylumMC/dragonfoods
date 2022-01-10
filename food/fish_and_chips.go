package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type FishAndChips struct{}

// AlwaysConsumable ...
func (FishAndChips) AlwaysConsumable() bool {
	return false
}

// Category ...
func (FishAndChips) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (FishAndChips) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(9, 5.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (FishAndChips) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (FishAndChips) Edible() bool {
	return true
}

// EncodeItem ...
func (FishAndChips) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:fish_and_chips", 0
}

// Name ...
func (FishAndChips) Name() string {
	return "FishAndChips"
}

// Texture ...
func (FishAndChips) Texture() image.Image {
	return textureFromName("fishandchips")
}
