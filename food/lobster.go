package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Lobster struct{}

// AlwaysConsumable ...
func (Lobster) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Lobster) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Lobster) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Lobster) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Lobster) Edible() bool {
	return true
}

// EncodeItem ...
func (Lobster) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:lobster", 0
}

// Name ...
func (Lobster) Name() string {
	return "Lobster"
}

// Texture ...
func (Lobster) Texture() image.Image {
	return textureFromName("lobster")
}
