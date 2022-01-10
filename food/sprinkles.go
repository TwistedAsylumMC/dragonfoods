package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Sprinkles struct{}

// AlwaysConsumable ...
func (Sprinkles) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Sprinkles) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Sprinkles) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Sprinkles) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (Sprinkles) Edible() bool {
	return true
}

// EncodeItem ...
func (Sprinkles) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sprinkles", 0
}

// Name ...
func (Sprinkles) Name() string {
	return "Sprinkles"
}

// Texture ...
func (Sprinkles) Texture() image.Image {
	return textureFromName("sprinkles")
}
