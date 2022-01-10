package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Gum struct{}

// AlwaysConsumable ...
func (Gum) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Gum) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Gum) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Gum) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Gum) Edible() bool {
	return true
}

// EncodeItem ...
func (Gum) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:gum", 0
}

// Name ...
func (Gum) Name() string {
	return "Gum"
}

// Texture ...
func (Gum) Texture() image.Image {
	return textureFromName("gum")
}
