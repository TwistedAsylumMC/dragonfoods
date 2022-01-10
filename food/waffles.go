package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Waffles struct{}

// AlwaysConsumable ...
func (Waffles) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Waffles) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Waffles) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Waffles) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Waffles) Edible() bool {
	return true
}

// EncodeItem ...
func (Waffles) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:waffles", 0
}

// Name ...
func (Waffles) Name() string {
	return "Waffles"
}

// Texture ...
func (Waffles) Texture() image.Image {
	return textureFromName("waffles")
}
