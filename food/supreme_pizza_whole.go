package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SupremePizzaWhole struct{}

// AlwaysConsumable ...
func (SupremePizzaWhole) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SupremePizzaWhole) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SupremePizzaWhole) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(18, 10.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (SupremePizzaWhole) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SupremePizzaWhole) Edible() bool {
	return true
}

// EncodeItem ...
func (SupremePizzaWhole) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:supreme_pizza_whole", 0
}

// Name ...
func (SupremePizzaWhole) Name() string {
	return "SupremePizzaWhole"
}

// Texture ...
func (SupremePizzaWhole) Texture() image.Image {
	return textureFromName("supreme1")
}
