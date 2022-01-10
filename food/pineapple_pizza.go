package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PineapplePizza struct{}

// AlwaysConsumable ...
func (PineapplePizza) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PineapplePizza) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PineapplePizza) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (PineapplePizza) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PineapplePizza) Edible() bool {
	return true
}

// EncodeItem ...
func (PineapplePizza) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pineapple_pizza", 0
}

// Name ...
func (PineapplePizza) Name() string {
	return "Pineapple Pizza"
}

// Texture ...
func (PineapplePizza) Texture() image.Image {
	return textureFromName("pizza5")
}
