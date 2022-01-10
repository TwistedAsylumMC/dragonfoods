package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type StickOfButter struct{}

// AlwaysConsumable ...
func (StickOfButter) AlwaysConsumable() bool {
	return false
}

// Category ...
func (StickOfButter) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (StickOfButter) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (StickOfButter) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (StickOfButter) Edible() bool {
	return true
}

// EncodeItem ...
func (StickOfButter) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:stick_of_butter", 0
}

// Name ...
func (StickOfButter) Name() string {
	return "StickOfButter"
}

// Texture ...
func (StickOfButter) Texture() image.Image {
	return textureFromName("stickofbutter")
}
