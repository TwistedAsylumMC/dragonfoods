package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChorizoPizzaWhole struct{}

// AlwaysConsumable ...
func (ChorizoPizzaWhole) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChorizoPizzaWhole) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChorizoPizzaWhole) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(19, 11.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChorizoPizzaWhole) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChorizoPizzaWhole) Edible() bool {
	return true
}

// EncodeItem ...
func (ChorizoPizzaWhole) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chorizo_pizza_whole", 0
}

// Name ...
func (ChorizoPizzaWhole) Name() string {
	return "ChorizoPizzaWhole"
}

// Texture ...
func (ChorizoPizzaWhole) Texture() image.Image {
	return textureFromName("bbq1")
}
