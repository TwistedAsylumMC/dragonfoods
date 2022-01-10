package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Pineapple struct{}

// AlwaysConsumable ...
func (Pineapple) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Pineapple) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Pineapple) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (Pineapple) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Pineapple) Edible() bool {
	return true
}

// EncodeItem ...
func (Pineapple) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pineapple", 0
}

// Name ...
func (Pineapple) Name() string {
	return "Pineapple"
}

// Texture ...
func (Pineapple) Texture() image.Image {
	return textureFromName("pineapple_item")
}
