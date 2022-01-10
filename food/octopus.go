package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Octopus struct{}

// AlwaysConsumable ...
func (Octopus) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Octopus) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Octopus) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Octopus) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Octopus) Edible() bool {
	return true
}

// EncodeItem ...
func (Octopus) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:octopus", 0
}

// Name ...
func (Octopus) Name() string {
	return "Octopus"
}

// Texture ...
func (Octopus) Texture() image.Image {
	return textureFromName("octopus")
}
