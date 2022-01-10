package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type VanillaYogurt struct{}

// AlwaysConsumable ...
func (VanillaYogurt) AlwaysConsumable() bool {
	return false
}

// Category ...
func (VanillaYogurt) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (VanillaYogurt) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (VanillaYogurt) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (VanillaYogurt) Edible() bool {
	return true
}

// EncodeItem ...
func (VanillaYogurt) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:vanilla_yogurt", 0
}

// Name ...
func (VanillaYogurt) Name() string {
	return "VanillaYogurt"
}

// Texture ...
func (VanillaYogurt) Texture() image.Image {
	return textureFromName("vanilla_or_lemon_ice_cream")
}
