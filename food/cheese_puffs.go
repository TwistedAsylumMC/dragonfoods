package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CheesePuffs struct{}

// AlwaysConsumable ...
func (CheesePuffs) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CheesePuffs) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CheesePuffs) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (CheesePuffs) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CheesePuffs) Edible() bool {
	return true
}

// EncodeItem ...
func (CheesePuffs) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cheese_puffs", 0
}

// Name ...
func (CheesePuffs) Name() string {
	return "Cheese Puffs"
}

// Texture ...
func (CheesePuffs) Texture() image.Image {
	return textureFromName("25_cheesepuff_bowl")
}
