package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Lettuce struct{}

// AlwaysConsumable ...
func (Lettuce) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Lettuce) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Lettuce) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Lettuce) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Lettuce) Edible() bool {
	return true
}

// EncodeItem ...
func (Lettuce) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:lettuce", 0
}

// Name ...
func (Lettuce) Name() string {
	return "Lettuce"
}

// Texture ...
func (Lettuce) Texture() image.Image {
	return textureFromName("lettuce_leaf")
}
