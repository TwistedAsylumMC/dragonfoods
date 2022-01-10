package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedMeatball struct{}

// AlwaysConsumable ...
func (CookedMeatball) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedMeatball) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedMeatball) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedMeatball) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedMeatball) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedMeatball) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_meatball", 0
}

// Name ...
func (CookedMeatball) Name() string {
	return "CookedMeatball"
}

// Texture ...
func (CookedMeatball) Texture() image.Image {
	return textureFromName("69_meatball")
}
