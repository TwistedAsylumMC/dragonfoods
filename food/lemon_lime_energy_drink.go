package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type LemonLimeEnergyDrink struct{}

// AlwaysConsumable ...
func (LemonLimeEnergyDrink) AlwaysConsumable() bool {
	return false
}

// Category ...
func (LemonLimeEnergyDrink) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (LemonLimeEnergyDrink) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (LemonLimeEnergyDrink) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (LemonLimeEnergyDrink) Edible() bool {
	return true
}

// EncodeItem ...
func (LemonLimeEnergyDrink) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:lemon_lime_energy_drink", 0
}

// Name ...
func (LemonLimeEnergyDrink) Name() string {
	return "Lemon Lime Energy Drink"
}

// Texture ...
func (LemonLimeEnergyDrink) Texture() image.Image {
	return textureFromName("energy-grun")
}
