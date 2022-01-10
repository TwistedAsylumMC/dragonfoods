package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Pancakes struct{}

// AlwaysConsumable ...
func (Pancakes) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Pancakes) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Pancakes) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Pancakes) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Pancakes) Edible() bool {
	return true
}

// EncodeItem ...
func (Pancakes) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pancakes", 0
}

// Name ...
func (Pancakes) Name() string {
	return "Pancakes"
}

// Texture ...
func (Pancakes) Texture() image.Image {
	return textureFromName("pancake")
}
