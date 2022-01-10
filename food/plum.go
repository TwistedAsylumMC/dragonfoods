package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Plum struct{}

// AlwaysConsumable ...
func (Plum) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Plum) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Plum) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Plum) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Plum) Edible() bool {
	return true
}

// EncodeItem ...
func (Plum) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:plum", 0
}

// Name ...
func (Plum) Name() string {
	return "Plum"
}

// Texture ...
func (Plum) Texture() image.Image {
	return textureFromName("plum")
}
