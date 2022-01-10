package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RootbeerFloat struct{}

// AlwaysConsumable ...
func (RootbeerFloat) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RootbeerFloat) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RootbeerFloat) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (RootbeerFloat) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RootbeerFloat) Edible() bool {
	return true
}

// EncodeItem ...
func (RootbeerFloat) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:rootbeer_float", 0
}

// Name ...
func (RootbeerFloat) Name() string {
	return "RootbeerFloat"
}

// Texture ...
func (RootbeerFloat) Texture() image.Image {
	return textureFromName("soda_glass")
}
