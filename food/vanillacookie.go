package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Vanillacookie struct{}

// AlwaysConsumable ...
func (Vanillacookie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Vanillacookie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Vanillacookie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Vanillacookie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Vanillacookie) Edible() bool {
	return true
}

// EncodeItem ...
func (Vanillacookie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:vanillacookie", 0
}

// Name ...
func (Vanillacookie) Name() string {
	return "Vanillacookie"
}

// Texture ...
func (Vanillacookie) Texture() image.Image {
	return textureFromName("cookie3")
}
