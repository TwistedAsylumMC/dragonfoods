package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BottleOfWhiteWine struct{}

// AlwaysConsumable ...
func (BottleOfWhiteWine) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BottleOfWhiteWine) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BottleOfWhiteWine) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BottleOfWhiteWine) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BottleOfWhiteWine) Edible() bool {
	return true
}

// EncodeItem ...
func (BottleOfWhiteWine) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bottle_of_white_wine", 0
}

// Name ...
func (BottleOfWhiteWine) Name() string {
	return "Bottle Of White Wine"
}

// Texture ...
func (BottleOfWhiteWine) Texture() image.Image {
	return textureFromName("wine_white")
}
