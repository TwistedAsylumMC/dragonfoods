package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type HamSub struct{}

// AlwaysConsumable ...
func (HamSub) AlwaysConsumable() bool {
	return false
}

// Category ...
func (HamSub) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (HamSub) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (HamSub) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (HamSub) Edible() bool {
	return true
}

// EncodeItem ...
func (HamSub) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:ham_sub", 0
}

// Name ...
func (HamSub) Name() string {
	return "Ham Sub"
}

// Texture ...
func (HamSub) Texture() image.Image {
	return textureFromName("sub2")
}
