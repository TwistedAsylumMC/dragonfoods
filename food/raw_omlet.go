package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RawOmlet struct{}

// AlwaysConsumable ...
func (RawOmlet) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RawOmlet) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RawOmlet) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (RawOmlet) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RawOmlet) Edible() bool {
	return true
}

// EncodeItem ...
func (RawOmlet) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:raw_omlet", 0
}

// Name ...
func (RawOmlet) Name() string {
	return "Raw Omlet"
}

// Texture ...
func (RawOmlet) Texture() image.Image {
	return textureFromName("73_omlet")
}
