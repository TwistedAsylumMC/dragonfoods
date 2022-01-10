package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Croissant struct{}

// AlwaysConsumable ...
func (Croissant) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Croissant) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Croissant) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Croissant) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Croissant) Edible() bool {
	return true
}

// EncodeItem ...
func (Croissant) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:croissant", 0
}

// Name ...
func (Croissant) Name() string {
	return "Croissant"
}

// Texture ...
func (Croissant) Texture() image.Image {
	return textureFromName("croissant")
}
