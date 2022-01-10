package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type HuevosRancheros struct{}

// AlwaysConsumable ...
func (HuevosRancheros) AlwaysConsumable() bool {
	return false
}

// Category ...
func (HuevosRancheros) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (HuevosRancheros) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (HuevosRancheros) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (HuevosRancheros) Edible() bool {
	return true
}

// EncodeItem ...
func (HuevosRancheros) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:huevos_rancheros", 0
}

// Name ...
func (HuevosRancheros) Name() string {
	return "Huevos Rancheros"
}

// Texture ...
func (HuevosRancheros) Texture() image.Image {
	return textureFromName("egg2")
}
