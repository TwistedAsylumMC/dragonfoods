package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Empanada struct{}

// AlwaysConsumable ...
func (Empanada) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Empanada) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Empanada) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (Empanada) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Empanada) Edible() bool {
	return true
}

// EncodeItem ...
func (Empanada) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:empanada", 0
}

// Name ...
func (Empanada) Name() string {
	return "Empanada"
}

// Texture ...
func (Empanada) Texture() image.Image {
	return textureFromName("empanada")
}
