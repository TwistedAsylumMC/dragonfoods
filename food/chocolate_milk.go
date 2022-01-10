package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChocolateMilk struct{}

// AlwaysConsumable ...
func (ChocolateMilk) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChocolateMilk) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChocolateMilk) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChocolateMilk) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChocolateMilk) Edible() bool {
	return true
}

// EncodeItem ...
func (ChocolateMilk) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chocolate_milk", 0
}

// Name ...
func (ChocolateMilk) Name() string {
	return "Chocolate Milk"
}

// Texture ...
func (ChocolateMilk) Texture() image.Image {
	return textureFromName("cup")
}
