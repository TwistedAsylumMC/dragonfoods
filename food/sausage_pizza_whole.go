package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SausagePizzaWhole struct{}

// AlwaysConsumable ...
func (SausagePizzaWhole) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SausagePizzaWhole) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SausagePizzaWhole) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(19, 11.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (SausagePizzaWhole) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SausagePizzaWhole) Edible() bool {
	return true
}

// EncodeItem ...
func (SausagePizzaWhole) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sausage_pizza_whole", 0
}

// Name ...
func (SausagePizzaWhole) Name() string {
	return "SausagePizzaWhole"
}

// Texture ...
func (SausagePizzaWhole) Texture() image.Image {
	return textureFromName("sausage1")
}
