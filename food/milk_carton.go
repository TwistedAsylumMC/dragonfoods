package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MilkCarton struct{}

// AlwaysConsumable ...
func (MilkCarton) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MilkCarton) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MilkCarton) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (MilkCarton) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MilkCarton) Edible() bool {
	return true
}

// EncodeItem ...
func (MilkCarton) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:milk_carton", 0
}

// Name ...
func (MilkCarton) Name() string {
	return "Milk Carton"
}

// Texture ...
func (MilkCarton) Texture() image.Image {
	return textureFromName("milk")
}
