package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BottleOfBlueWine struct{}

// AlwaysConsumable ...
func (BottleOfBlueWine) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BottleOfBlueWine) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BottleOfBlueWine) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(11, 6.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BottleOfBlueWine) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BottleOfBlueWine) Edible() bool {
	return true
}

// EncodeItem ...
func (BottleOfBlueWine) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bottle_of_blue_wine", 0
}

// Name ...
func (BottleOfBlueWine) Name() string {
	return "Bottle Of Blue Wine"
}

// Texture ...
func (BottleOfBlueWine) Texture() image.Image {
	return textureFromName("wine_white2")
}
