package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RawMeatball struct{}

// AlwaysConsumable ...
func (RawMeatball) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RawMeatball) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RawMeatball) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (RawMeatball) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RawMeatball) Edible() bool {
	return true
}

// EncodeItem ...
func (RawMeatball) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:raw_meatball", 0
}

// Name ...
func (RawMeatball) Name() string {
	return "Raw Meatball"
}

// Texture ...
func (RawMeatball) Texture() image.Image {
	return textureFromName("meatball")
}
