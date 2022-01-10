package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type EnderCola struct{}

// AlwaysConsumable ...
func (EnderCola) AlwaysConsumable() bool {
	return false
}

// Category ...
func (EnderCola) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (EnderCola) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (EnderCola) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (EnderCola) Edible() bool {
	return true
}

// EncodeItem ...
func (EnderCola) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:ender_cola", 0
}

// Name ...
func (EnderCola) Name() string {
	return "Ender Cola"
}

// Texture ...
func (EnderCola) Texture() image.Image {
	return textureFromName("grape_soda")
}
