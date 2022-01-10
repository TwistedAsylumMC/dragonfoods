package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Mustard struct{}

// AlwaysConsumable ...
func (Mustard) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Mustard) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Mustard) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Mustard) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Mustard) Edible() bool {
	return true
}

// EncodeItem ...
func (Mustard) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:mustard", 0
}

// Name ...
func (Mustard) Name() string {
	return "Mustard"
}

// Texture ...
func (Mustard) Texture() image.Image {
	return textureFromName("mustard")
}
