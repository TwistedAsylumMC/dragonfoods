package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Pizza struct{}

// AlwaysConsumable ...
func (Pizza) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Pizza) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Pizza) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Pizza) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Pizza) Edible() bool {
	return true
}

// EncodeItem ...
func (Pizza) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pizza", 0
}

// Name ...
func (Pizza) Name() string {
	return "Pizza"
}

// Texture ...
func (Pizza) Texture() image.Image {
	return textureFromName("pizza_02")
}
