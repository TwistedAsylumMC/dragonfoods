package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Baguette struct{}

// AlwaysConsumable ...
func (Baguette) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Baguette) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Baguette) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Baguette) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Baguette) Edible() bool {
	return true
}

// EncodeItem ...
func (Baguette) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:baguette", 0
}

// Name ...
func (Baguette) Name() string {
	return "Baguette"
}

// Texture ...
func (Baguette) Texture() image.Image {
	return textureFromName("baguette")
}
