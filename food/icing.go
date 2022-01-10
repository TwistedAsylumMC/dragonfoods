package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Icing struct{}

// AlwaysConsumable ...
func (Icing) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Icing) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Icing) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Icing) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (Icing) Edible() bool {
	return true
}

// EncodeItem ...
func (Icing) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:icing", 0
}

// Name ...
func (Icing) Name() string {
	return "Icing"
}

// Texture ...
func (Icing) Texture() image.Image {
	return textureFromName("icing")
}
