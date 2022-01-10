package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type EnergyBar struct{}

// AlwaysConsumable ...
func (EnergyBar) AlwaysConsumable() bool {
	return false
}

// Category ...
func (EnergyBar) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (EnergyBar) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (EnergyBar) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (EnergyBar) Edible() bool {
	return true
}

// EncodeItem ...
func (EnergyBar) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:energy_bar", 0
}

// Name ...
func (EnergyBar) Name() string {
	return "EnergyBar"
}

// Texture ...
func (EnergyBar) Texture() image.Image {
	return textureFromName("energy_bar")
}
