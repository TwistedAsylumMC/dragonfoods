package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChocolateMilkCarton struct{}

// AlwaysConsumable ...
func (ChocolateMilkCarton) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChocolateMilkCarton) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChocolateMilkCarton) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChocolateMilkCarton) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChocolateMilkCarton) Edible() bool {
	return true
}

// EncodeItem ...
func (ChocolateMilkCarton) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chocolate_milk_carton", 0
}

// Name ...
func (ChocolateMilkCarton) Name() string {
	return "Chocolate Milk Carton"
}

// Texture ...
func (ChocolateMilkCarton) Texture() image.Image {
	return textureFromName("milk2")
}
