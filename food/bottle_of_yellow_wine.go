package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BottleOfYellowWine struct{}

// AlwaysConsumable ...
func (BottleOfYellowWine) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BottleOfYellowWine) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BottleOfYellowWine) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(11, 6.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BottleOfYellowWine) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BottleOfYellowWine) Edible() bool {
	return true
}

// EncodeItem ...
func (BottleOfYellowWine) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bottle_of_yellow_wine", 0
}

// Name ...
func (BottleOfYellowWine) Name() string {
	return "Bottle Of Yellow Wine"
}

// Texture ...
func (BottleOfYellowWine) Texture() image.Image {
	return textureFromName("wine_white3")
}
