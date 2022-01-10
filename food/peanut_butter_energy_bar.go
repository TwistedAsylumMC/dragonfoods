package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PeanutButterEnergyBar struct{}

// AlwaysConsumable ...
func (PeanutButterEnergyBar) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PeanutButterEnergyBar) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PeanutButterEnergyBar) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (PeanutButterEnergyBar) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PeanutButterEnergyBar) Edible() bool {
	return true
}

// EncodeItem ...
func (PeanutButterEnergyBar) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:peanut_butter_energy_bar", 0
}

// Name ...
func (PeanutButterEnergyBar) Name() string {
	return "Peanut Butter Energy Bar"
}

// Texture ...
func (PeanutButterEnergyBar) Texture() image.Image {
	return textureFromName("energy_bar2")
}
