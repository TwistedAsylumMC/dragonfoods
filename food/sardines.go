package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Sardines struct{}

// AlwaysConsumable ...
func (Sardines) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Sardines) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Sardines) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Sardines) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Sardines) Edible() bool {
	return true
}

// EncodeItem ...
func (Sardines) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sardines", 0
}

// Name ...
func (Sardines) Name() string {
	return "Sardines"
}

// Texture ...
func (Sardines) Texture() image.Image {
	return textureFromName("sardines")
}
