package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Salt struct{}

// AlwaysConsumable ...
func (Salt) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Salt) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Salt) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(1, 0.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Salt) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Salt) Edible() bool {
	return true
}

// EncodeItem ...
func (Salt) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:salt", 0
}

// Name ...
func (Salt) Name() string {
	return "Salt"
}

// Texture ...
func (Salt) Texture() image.Image {
	return textureFromName("salt")
}
