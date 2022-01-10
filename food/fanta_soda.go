package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type FantaSoda struct{}

// AlwaysConsumable ...
func (FantaSoda) AlwaysConsumable() bool {
	return false
}

// Category ...
func (FantaSoda) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (FantaSoda) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (FantaSoda) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (FantaSoda) Edible() bool {
	return true
}

// EncodeItem ...
func (FantaSoda) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:fanta_soda", 0
}

// Name ...
func (FantaSoda) Name() string {
	return "Fanta Soda"
}

// Texture ...
func (FantaSoda) Texture() image.Image {
	return textureFromName("sodabottel2")
}
