package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Blt struct{}

// AlwaysConsumable ...
func (Blt) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Blt) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Blt) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (Blt) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Blt) Edible() bool {
	return true
}

// EncodeItem ...
func (Blt) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blt", 0
}

// Name ...
func (Blt) Name() string {
	return "Blt"
}

// Texture ...
func (Blt) Texture() image.Image {
	return textureFromName("blt")
}
