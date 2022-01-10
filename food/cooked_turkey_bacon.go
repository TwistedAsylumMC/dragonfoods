package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedTurkeyBacon struct{}

// AlwaysConsumable ...
func (CookedTurkeyBacon) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedTurkeyBacon) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedTurkeyBacon) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedTurkeyBacon) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedTurkeyBacon) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedTurkeyBacon) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_turkey_bacon", 0
}

// Name ...
func (CookedTurkeyBacon) Name() string {
	return "Cooked Turkey Bacon"
}

// Texture ...
func (CookedTurkeyBacon) Texture() image.Image {
	return textureFromName("cookedtbacon")
}
