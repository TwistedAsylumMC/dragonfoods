package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BakedBiscut struct{}

// AlwaysConsumable ...
func (BakedBiscut) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BakedBiscut) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BakedBiscut) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BakedBiscut) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BakedBiscut) Edible() bool {
	return true
}

// EncodeItem ...
func (BakedBiscut) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:baked_biscut", 0
}

// Name ...
func (BakedBiscut) Name() string {
	return "BakedBiscut"
}

// Texture ...
func (BakedBiscut) Texture() image.Image {
	return textureFromName("bakedbiscut")
}
