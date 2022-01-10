package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Sandwich struct{}

// AlwaysConsumable ...
func (Sandwich) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Sandwich) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Sandwich) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Sandwich) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Sandwich) Edible() bool {
	return true
}

// EncodeItem ...
func (Sandwich) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sandwich", 0
}

// Name ...
func (Sandwich) Name() string {
	return "Sandwich"
}

// Texture ...
func (Sandwich) Texture() image.Image {
	return textureFromName("sandwich")
}
