package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Steak struct{}

// AlwaysConsumable ...
func (Steak) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Steak) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Steak) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Steak) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Steak) Edible() bool {
	return true
}

// EncodeItem ...
func (Steak) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:steak", 0
}

// Name ...
func (Steak) Name() string {
	return "Steak"
}

// Texture ...
func (Steak) Texture() image.Image {
	return textureFromName("steak")
}
