package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BottleOfRedWine struct{}

// AlwaysConsumable ...
func (BottleOfRedWine) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BottleOfRedWine) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BottleOfRedWine) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(11, 6.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BottleOfRedWine) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BottleOfRedWine) Edible() bool {
	return true
}

// EncodeItem ...
func (BottleOfRedWine) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bottle_of_red_wine", 0
}

// Name ...
func (BottleOfRedWine) Name() string {
	return "BottleOfRedWine"
}

// Texture ...
func (BottleOfRedWine) Texture() image.Image {
	return textureFromName("wine_red")
}
