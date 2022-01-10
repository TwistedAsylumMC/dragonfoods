package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Hops struct{}

// AlwaysConsumable ...
func (Hops) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Hops) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Hops) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Hops) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Hops) Edible() bool {
	return true
}

// EncodeItem ...
func (Hops) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:hops", 0
}

// Name ...
func (Hops) Name() string {
	return "Hops"
}

// Texture ...
func (Hops) Texture() image.Image {
	return textureFromName("hops_01")
}
