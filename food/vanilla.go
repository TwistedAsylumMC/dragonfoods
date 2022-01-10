package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Vanilla struct{}

// AlwaysConsumable ...
func (Vanilla) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Vanilla) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Vanilla) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Vanilla) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (Vanilla) Edible() bool {
	return true
}

// EncodeItem ...
func (Vanilla) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:vanilla", 0
}

// Name ...
func (Vanilla) Name() string {
	return "Vanilla"
}

// Texture ...
func (Vanilla) Texture() image.Image {
	return textureFromName("vanilla")
}
