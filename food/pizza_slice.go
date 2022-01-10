package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PizzaSlice struct{}

// AlwaysConsumable ...
func (PizzaSlice) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PizzaSlice) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PizzaSlice) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (PizzaSlice) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PizzaSlice) Edible() bool {
	return true
}

// EncodeItem ...
func (PizzaSlice) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pizza_slice", 0
}

// Name ...
func (PizzaSlice) Name() string {
	return "Pizza Slice"
}

// Texture ...
func (PizzaSlice) Texture() image.Image {
	return textureFromName("pizza_1")
}
