package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Marmalade struct{}

// AlwaysConsumable ...
func (Marmalade) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Marmalade) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Marmalade) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Marmalade) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Marmalade) Edible() bool {
	return true
}

// EncodeItem ...
func (Marmalade) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:marmalade", 0
}

// Name ...
func (Marmalade) Name() string {
	return "Marmalade"
}

// Texture ...
func (Marmalade) Texture() image.Image {
	return textureFromName("marmalade")
}
