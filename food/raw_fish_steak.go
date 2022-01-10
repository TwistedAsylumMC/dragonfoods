package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RawFishSteak struct{}

// AlwaysConsumable ...
func (RawFishSteak) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RawFishSteak) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RawFishSteak) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (RawFishSteak) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RawFishSteak) Edible() bool {
	return true
}

// EncodeItem ...
func (RawFishSteak) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:raw_fish_steak", 0
}

// Name ...
func (RawFishSteak) Name() string {
	return "Raw Fish Steak"
}

// Texture ...
func (RawFishSteak) Texture() image.Image {
	return textureFromName("fishsteak")
}
