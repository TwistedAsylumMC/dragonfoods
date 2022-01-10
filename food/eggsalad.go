package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Eggsalad struct{}

// AlwaysConsumable ...
func (Eggsalad) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Eggsalad) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Eggsalad) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(11, 6.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Eggsalad) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Eggsalad) Edible() bool {
	return true
}

// EncodeItem ...
func (Eggsalad) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:eggsalad", 0
}

// Name ...
func (Eggsalad) Name() string {
	return "Eggsalad"
}

// Texture ...
func (Eggsalad) Texture() image.Image {
	return textureFromName("41_eggsalad_bowl")
}
