package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedBacon struct{}

// AlwaysConsumable ...
func (CookedBacon) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedBacon) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedBacon) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(3, 1.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedBacon) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedBacon) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedBacon) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_bacon", 0
}

// Name ...
func (CookedBacon) Name() string {
	return "CookedBacon"
}

// Texture ...
func (CookedBacon) Texture() image.Image {
	return textureFromName("13_bacon")
}
