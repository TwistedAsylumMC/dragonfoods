package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Shrimp struct{}

// AlwaysConsumable ...
func (Shrimp) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Shrimp) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Shrimp) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Shrimp) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Shrimp) Edible() bool {
	return true
}

// EncodeItem ...
func (Shrimp) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:shrimp", 0
}

// Name ...
func (Shrimp) Name() string {
	return "Shrimp"
}

// Texture ...
func (Shrimp) Texture() image.Image {
	return textureFromName("shrimp2")
}
