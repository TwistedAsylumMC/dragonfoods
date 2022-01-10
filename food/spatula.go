package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Spatula struct{}

// AlwaysConsumable ...
func (Spatula) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Spatula) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Spatula) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Spatula) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (Spatula) Edible() bool {
	return true
}

// EncodeItem ...
func (Spatula) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:spatula", 0
}

// Name ...
func (Spatula) Name() string {
	return "Spatula"
}

// Texture ...
func (Spatula) Texture() image.Image {
	return textureFromName("spatula")
}
