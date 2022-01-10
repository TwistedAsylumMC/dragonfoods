package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Onion struct{}

// AlwaysConsumable ...
func (Onion) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Onion) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Onion) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Onion) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Onion) Edible() bool {
	return true
}

// EncodeItem ...
func (Onion) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:onion", 0
}

// Name ...
func (Onion) Name() string {
	return "Onion"
}

// Texture ...
func (Onion) Texture() image.Image {
	return textureFromName("onion_cebolla")
}
