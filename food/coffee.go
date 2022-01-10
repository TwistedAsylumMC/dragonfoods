package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Coffee struct{}

// AlwaysConsumable ...
func (Coffee) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Coffee) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Coffee) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Coffee) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Coffee) Edible() bool {
	return true
}

// EncodeItem ...
func (Coffee) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:coffee", 0
}

// Name ...
func (Coffee) Name() string {
	return "Coffee"
}

// Texture ...
func (Coffee) Texture() image.Image {
	return textureFromName("coffee_mug")
}
