package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Burger struct{}

// AlwaysConsumable ...
func (Burger) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Burger) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Burger) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Burger) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Burger) Edible() bool {
	return true
}

// EncodeItem ...
func (Burger) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:burger", 0
}

// Name ...
func (Burger) Name() string {
	return "Burger"
}

// Texture ...
func (Burger) Texture() image.Image {
	return textureFromName("15_burger")
}
