package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type EnderColaZero struct{}

// AlwaysConsumable ...
func (EnderColaZero) AlwaysConsumable() bool {
	return false
}

// Category ...
func (EnderColaZero) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (EnderColaZero) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (EnderColaZero) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (EnderColaZero) Edible() bool {
	return true
}

// EncodeItem ...
func (EnderColaZero) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:ender_cola_zero", 0
}

// Name ...
func (EnderColaZero) Name() string {
	return "EnderColaZero"
}

// Texture ...
func (EnderColaZero) Texture() image.Image {
	return textureFromName("grape_soda2")
}
