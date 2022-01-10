package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Leek struct{}

// AlwaysConsumable ...
func (Leek) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Leek) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Leek) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Leek) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Leek) Edible() bool {
	return true
}

// EncodeItem ...
func (Leek) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:leek", 0
}

// Name ...
func (Leek) Name() string {
	return "Leek"
}

// Texture ...
func (Leek) Texture() image.Image {
	return textureFromName("leek")
}
