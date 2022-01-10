package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Fig struct{}

// AlwaysConsumable ...
func (Fig) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Fig) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Fig) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Fig) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Fig) Edible() bool {
	return true
}

// EncodeItem ...
func (Fig) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:fig", 0
}

// Name ...
func (Fig) Name() string {
	return "Fig"
}

// Texture ...
func (Fig) Texture() image.Image {
	return textureFromName("fig_01")
}
