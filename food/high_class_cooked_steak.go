package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type HighClassCookedSteak struct{}

// AlwaysConsumable ...
func (HighClassCookedSteak) AlwaysConsumable() bool {
	return false
}

// Category ...
func (HighClassCookedSteak) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (HighClassCookedSteak) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(15, 9.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (HighClassCookedSteak) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (HighClassCookedSteak) Edible() bool {
	return true
}

// EncodeItem ...
func (HighClassCookedSteak) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:high_class_cooked_steak", 0
}

// Name ...
func (HighClassCookedSteak) Name() string {
	return "High Class Cooked Steak"
}

// Texture ...
func (HighClassCookedSteak) Texture() image.Image {
	return textureFromName("96_steak_dish")
}
