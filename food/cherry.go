package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Cherry struct{}

// AlwaysConsumable ...
func (Cherry) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Cherry) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Cherry) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Cherry) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Cherry) Edible() bool {
	return true
}

// EncodeItem ...
func (Cherry) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cherry", 0
}

// Name ...
func (Cherry) Name() string {
	return "Cherry"
}

// Texture ...
func (Cherry) Texture() image.Image {
	return textureFromName("cherry")
}
