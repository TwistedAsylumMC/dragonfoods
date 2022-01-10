package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedFishFillet struct{}

// AlwaysConsumable ...
func (CookedFishFillet) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedFishFillet) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedFishFillet) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedFishFillet) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedFishFillet) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedFishFillet) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_fish_fillet", 0
}

// Name ...
func (CookedFishFillet) Name() string {
	return "Cooked Fish Fillet"
}

// Texture ...
func (CookedFishFillet) Texture() image.Image {
	return textureFromName("fishfillet2")
}
