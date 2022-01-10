package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SpicyCheesePuffs struct{}

// AlwaysConsumable ...
func (SpicyCheesePuffs) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SpicyCheesePuffs) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SpicyCheesePuffs) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(11, 6.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (SpicyCheesePuffs) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SpicyCheesePuffs) Edible() bool {
	return true
}

// EncodeItem ...
func (SpicyCheesePuffs) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:spicy_cheese_puffs", 0
}

// Name ...
func (SpicyCheesePuffs) Name() string {
	return "Spicy Cheese Puffs"
}

// Texture ...
func (SpicyCheesePuffs) Texture() image.Image {
	return textureFromName("26cheesepuff")
}
