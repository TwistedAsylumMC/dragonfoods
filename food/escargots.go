package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Escargots struct{}

// AlwaysConsumable ...
func (Escargots) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Escargots) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Escargots) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Escargots) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Escargots) Edible() bool {
	return true
}

// EncodeItem ...
func (Escargots) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:escargots", 0
}

// Name ...
func (Escargots) Name() string {
	return "Escargots"
}

// Texture ...
func (Escargots) Texture() image.Image {
	return textureFromName("escargots")
}
