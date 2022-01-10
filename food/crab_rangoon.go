package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CrabRangoon struct{}

// AlwaysConsumable ...
func (CrabRangoon) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CrabRangoon) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CrabRangoon) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (CrabRangoon) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CrabRangoon) Edible() bool {
	return true
}

// EncodeItem ...
func (CrabRangoon) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:crab_rangoon", 0
}

// Name ...
func (CrabRangoon) Name() string {
	return "CrabRangoon"
}

// Texture ...
func (CrabRangoon) Texture() image.Image {
	return textureFromName("crab1")
}
