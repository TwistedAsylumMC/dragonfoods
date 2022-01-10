package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Bun struct{}

// AlwaysConsumable ...
func (Bun) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Bun) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Bun) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(1, 0.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Bun) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Bun) Edible() bool {
	return true
}

// EncodeItem ...
func (Bun) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bun", 0
}

// Name ...
func (Bun) Name() string {
	return "Bun"
}

// Texture ...
func (Bun) Texture() image.Image {
	return textureFromName("bun")
}
