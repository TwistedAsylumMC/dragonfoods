package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Tuna struct{}

// AlwaysConsumable ...
func (Tuna) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Tuna) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Tuna) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (Tuna) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Tuna) Edible() bool {
	return true
}

// EncodeItem ...
func (Tuna) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:tuna", 0
}

// Name ...
func (Tuna) Name() string {
	return "Tuna"
}

// Texture ...
func (Tuna) Texture() image.Image {
	return textureFromName("tuna_can")
}
