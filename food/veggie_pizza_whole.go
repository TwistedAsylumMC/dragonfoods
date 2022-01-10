package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type VeggiePizzaWhole struct{}

// AlwaysConsumable ...
func (VeggiePizzaWhole) AlwaysConsumable() bool {
	return false
}

// Category ...
func (VeggiePizzaWhole) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (VeggiePizzaWhole) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(20, 12.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (VeggiePizzaWhole) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (VeggiePizzaWhole) Edible() bool {
	return true
}

// EncodeItem ...
func (VeggiePizzaWhole) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:veggie_pizza_whole", 0
}

// Name ...
func (VeggiePizzaWhole) Name() string {
	return "VeggiePizzaWhole"
}

// Texture ...
func (VeggiePizzaWhole) Texture() image.Image {
	return textureFromName("veggie1")
}
