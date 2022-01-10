package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Radish struct{}

// AlwaysConsumable ...
func (Radish) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Radish) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Radish) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(3, 1.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (Radish) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Radish) Edible() bool {
	return true
}

// EncodeItem ...
func (Radish) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:radish", 0
}

// Name ...
func (Radish) Name() string {
	return "Radish"
}

// Texture ...
func (Radish) Texture() image.Image {
	return textureFromName("radish")
}
