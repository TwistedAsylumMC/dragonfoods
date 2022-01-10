package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type UncookedBacon struct{}

// AlwaysConsumable ...
func (UncookedBacon) AlwaysConsumable() bool {
	return false
}

// Category ...
func (UncookedBacon) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (UncookedBacon) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(1, 0.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (UncookedBacon) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (UncookedBacon) Edible() bool {
	return true
}

// EncodeItem ...
func (UncookedBacon) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:uncooked_bacon", 0
}

// Name ...
func (UncookedBacon) Name() string {
	return "Uncooked Bacon"
}

// Texture ...
func (UncookedBacon) Texture() image.Image {
	return textureFromName("uncookedbacon13")
}
