package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Ham struct{}

// AlwaysConsumable ...
func (Ham) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Ham) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Ham) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(11, 6.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Ham) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Ham) Edible() bool {
	return true
}

// EncodeItem ...
func (Ham) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:ham", 0
}

// Name ...
func (Ham) Name() string {
	return "Ham"
}

// Texture ...
func (Ham) Texture() image.Image {
	return textureFromName("ham")
}
