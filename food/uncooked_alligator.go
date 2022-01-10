package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type UncookedAlligator struct{}

// AlwaysConsumable ...
func (UncookedAlligator) AlwaysConsumable() bool {
	return false
}

// Category ...
func (UncookedAlligator) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (UncookedAlligator) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (UncookedAlligator) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (UncookedAlligator) Edible() bool {
	return true
}

// EncodeItem ...
func (UncookedAlligator) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:uncooked_alligator", 0
}

// Name ...
func (UncookedAlligator) Name() string {
	return "UncookedAlligator"
}

// Texture ...
func (UncookedAlligator) Texture() image.Image {
	return textureFromName("alligator")
}
