package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedShrimp struct{}

// AlwaysConsumable ...
func (CookedShrimp) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedShrimp) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedShrimp) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedShrimp) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedShrimp) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedShrimp) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_shrimp", 0
}

// Name ...
func (CookedShrimp) Name() string {
	return "CookedShrimp"
}

// Texture ...
func (CookedShrimp) Texture() image.Image {
	return textureFromName("shrimp3")
}
