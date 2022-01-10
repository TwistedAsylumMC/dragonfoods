package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PeanutButter struct{}

// AlwaysConsumable ...
func (PeanutButter) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PeanutButter) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PeanutButter) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (PeanutButter) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PeanutButter) Edible() bool {
	return true
}

// EncodeItem ...
func (PeanutButter) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:peanut_butter", 0
}

// Name ...
func (PeanutButter) Name() string {
	return "PeanutButter"
}

// Texture ...
func (PeanutButter) Texture() image.Image {
	return textureFromName("peanut_butter")
}
