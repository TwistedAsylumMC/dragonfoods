package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type UncookedTurkeyBacon struct{}

// AlwaysConsumable ...
func (UncookedTurkeyBacon) AlwaysConsumable() bool {
	return false
}

// Category ...
func (UncookedTurkeyBacon) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (UncookedTurkeyBacon) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (UncookedTurkeyBacon) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (UncookedTurkeyBacon) Edible() bool {
	return true
}

// EncodeItem ...
func (UncookedTurkeyBacon) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:uncooked_turkey_bacon", 0
}

// Name ...
func (UncookedTurkeyBacon) Name() string {
	return "Uncooked Turkey Bacon"
}

// Texture ...
func (UncookedTurkeyBacon) Texture() image.Image {
	return textureFromName("bacon")
}
