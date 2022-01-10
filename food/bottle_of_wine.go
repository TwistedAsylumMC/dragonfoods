package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BottleOfWine struct{}

// AlwaysConsumable ...
func (BottleOfWine) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BottleOfWine) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BottleOfWine) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BottleOfWine) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BottleOfWine) Edible() bool {
	return true
}

// EncodeItem ...
func (BottleOfWine) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bottle_of_wine", 0
}

// Name ...
func (BottleOfWine) Name() string {
	return "Bottle Of Wine"
}

// Texture ...
func (BottleOfWine) Texture() image.Image {
	return textureFromName("winbottle")
}
