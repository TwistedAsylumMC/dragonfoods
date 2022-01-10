package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Loaf struct{}

// AlwaysConsumable ...
func (Loaf) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Loaf) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Loaf) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Loaf) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Loaf) Edible() bool {
	return true
}

// EncodeItem ...
func (Loaf) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:loaf", 0
}

// Name ...
func (Loaf) Name() string {
	return "Loaf"
}

// Texture ...
func (Loaf) Texture() image.Image {
	return textureFromName("loaf2")
}
