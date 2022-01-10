package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RawFishFillet struct{}

// AlwaysConsumable ...
func (RawFishFillet) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RawFishFillet) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RawFishFillet) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (RawFishFillet) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RawFishFillet) Edible() bool {
	return true
}

// EncodeItem ...
func (RawFishFillet) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:raw_fish_fillet", 0
}

// Name ...
func (RawFishFillet) Name() string {
	return "Raw Fish Fillet"
}

// Texture ...
func (RawFishFillet) Texture() image.Image {
	return textureFromName("fishfillet")
}
