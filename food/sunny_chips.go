package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SunnyChips struct{}

// AlwaysConsumable ...
func (SunnyChips) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SunnyChips) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SunnyChips) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (SunnyChips) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SunnyChips) Edible() bool {
	return true
}

// EncodeItem ...
func (SunnyChips) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sunny_chips", 0
}

// Name ...
func (SunnyChips) Name() string {
	return "SunnyChips"
}

// Texture ...
func (SunnyChips) Texture() image.Image {
	return textureFromName("snack1")
}
