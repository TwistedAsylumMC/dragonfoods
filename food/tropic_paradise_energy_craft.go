package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type TropicParadiseEnergyCraft struct{}

// AlwaysConsumable ...
func (TropicParadiseEnergyCraft) AlwaysConsumable() bool {
	return false
}

// Category ...
func (TropicParadiseEnergyCraft) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (TropicParadiseEnergyCraft) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (TropicParadiseEnergyCraft) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (TropicParadiseEnergyCraft) Edible() bool {
	return true
}

// EncodeItem ...
func (TropicParadiseEnergyCraft) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:tropic_paradise_energy_craft", 0
}

// Name ...
func (TropicParadiseEnergyCraft) Name() string {
	return "TropicParadiseEnergyCraft"
}

// Texture ...
func (TropicParadiseEnergyCraft) Texture() image.Image {
	return textureFromName("energy-lila")
}
