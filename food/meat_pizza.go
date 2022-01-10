package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MeatPizza struct{}

// AlwaysConsumable ...
func (MeatPizza) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MeatPizza) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MeatPizza) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (MeatPizza) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MeatPizza) Edible() bool {
	return true
}

// EncodeItem ...
func (MeatPizza) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:meat_pizza", 0
}

// Name ...
func (MeatPizza) Name() string {
	return "MeatPizza"
}

// Texture ...
func (MeatPizza) Texture() image.Image {
	return textureFromName("pizza4")
}
