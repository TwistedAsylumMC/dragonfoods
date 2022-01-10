package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Squash struct{}

// AlwaysConsumable ...
func (Squash) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Squash) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Squash) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Squash) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Squash) Edible() bool {
	return true
}

// EncodeItem ...
func (Squash) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:squash", 0
}

// Name ...
func (Squash) Name() string {
	return "Squash"
}

// Texture ...
func (Squash) Texture() image.Image {
	return textureFromName("squash_butternut_01")
}
