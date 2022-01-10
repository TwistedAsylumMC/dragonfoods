package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ArticFreezeEnergyDrink struct{}

// AlwaysConsumable ...
func (ArticFreezeEnergyDrink) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ArticFreezeEnergyDrink) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ArticFreezeEnergyDrink) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (ArticFreezeEnergyDrink) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ArticFreezeEnergyDrink) Edible() bool {
	return true
}

// EncodeItem ...
func (ArticFreezeEnergyDrink) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:artic_freeze_energy_drink", 0
}

// Name ...
func (ArticFreezeEnergyDrink) Name() string {
	return "ArticFreezeEnergyDrink"
}

// Texture ...
func (ArticFreezeEnergyDrink) Texture() image.Image {
	return textureFromName("energy-blau")
}
