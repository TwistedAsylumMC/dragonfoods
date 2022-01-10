package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedAlligator struct{}

// AlwaysConsumable ...
func (CookedAlligator) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedAlligator) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedAlligator) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(15, 9.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedAlligator) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedAlligator) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedAlligator) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_alligator", 0
}

// Name ...
func (CookedAlligator) Name() string {
	return "CookedAlligator"
}

// Texture ...
func (CookedAlligator) Texture() image.Image {
	return textureFromName("alligator2")
}
