package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedSteak struct{}

// AlwaysConsumable ...
func (CookedSteak) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedSteak) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedSteak) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedSteak) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedSteak) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedSteak) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_steak", 0
}

// Name ...
func (CookedSteak) Name() string {
	return "Cooked Steak"
}

// Texture ...
func (CookedSteak) Texture() image.Image {
	return textureFromName("95_steak")
}
