package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedOctopus struct{}

// AlwaysConsumable ...
func (CookedOctopus) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedOctopus) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedOctopus) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedOctopus) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedOctopus) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedOctopus) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_octopus", 0
}

// Name ...
func (CookedOctopus) Name() string {
	return "CookedOctopus"
}

// Texture ...
func (CookedOctopus) Texture() image.Image {
	return textureFromName("octopus2")
}
