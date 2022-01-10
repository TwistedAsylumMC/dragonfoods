package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Flour struct{}

// AlwaysConsumable ...
func (Flour) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Flour) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Flour) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Flour) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Flour) Edible() bool {
	return true
}

// EncodeItem ...
func (Flour) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:flour", 0
}

// Name ...
func (Flour) Name() string {
	return "Flour"
}

// Texture ...
func (Flour) Texture() image.Image {
	return textureFromName("flour")
}
