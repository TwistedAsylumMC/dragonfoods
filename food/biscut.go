package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Biscut struct{}

// AlwaysConsumable ...
func (Biscut) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Biscut) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Biscut) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Biscut) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Biscut) Edible() bool {
	return true
}

// EncodeItem ...
func (Biscut) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:biscut", 0
}

// Name ...
func (Biscut) Name() string {
	return "Biscut"
}

// Texture ...
func (Biscut) Texture() image.Image {
	return textureFromName("biscut")
}
