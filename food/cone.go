package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Cone struct{}

// AlwaysConsumable ...
func (Cone) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Cone) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Cone) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Cone) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Cone) Edible() bool {
	return true
}

// EncodeItem ...
func (Cone) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cone", 0
}

// Name ...
func (Cone) Name() string {
	return "Cone"
}

// Texture ...
func (Cone) Texture() image.Image {
	return textureFromName("cone")
}
