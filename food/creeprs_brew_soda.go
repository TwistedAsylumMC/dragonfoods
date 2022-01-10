package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CreeprsBrewSoda struct{}

// AlwaysConsumable ...
func (CreeprsBrewSoda) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CreeprsBrewSoda) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CreeprsBrewSoda) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (CreeprsBrewSoda) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CreeprsBrewSoda) Edible() bool {
	return true
}

// EncodeItem ...
func (CreeprsBrewSoda) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:creeprs_brew_soda", 0
}

// Name ...
func (CreeprsBrewSoda) Name() string {
	return "CreeprsBrewSoda"
}

// Texture ...
func (CreeprsBrewSoda) Texture() image.Image {
	return textureFromName("soft_drink_green")
}
