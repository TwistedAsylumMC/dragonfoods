package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RawRibs struct{}

// AlwaysConsumable ...
func (RawRibs) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RawRibs) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RawRibs) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (RawRibs) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RawRibs) Edible() bool {
	return true
}

// EncodeItem ...
func (RawRibs) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:raw_ribs", 0
}

// Name ...
func (RawRibs) Name() string {
	return "RawRibs"
}

// Texture ...
func (RawRibs) Texture() image.Image {
	return textureFromName("rawribs")
}
