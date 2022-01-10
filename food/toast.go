package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Toast struct{}

// AlwaysConsumable ...
func (Toast) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Toast) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Toast) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Toast) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Toast) Edible() bool {
	return true
}

// EncodeItem ...
func (Toast) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:toast", 0
}

// Name ...
func (Toast) Name() string {
	return "Toast"
}

// Texture ...
func (Toast) Texture() image.Image {
	return textureFromName("loaf")
}
