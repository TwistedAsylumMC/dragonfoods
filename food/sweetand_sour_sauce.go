package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SweetandSourSauce struct{}

// AlwaysConsumable ...
func (SweetandSourSauce) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SweetandSourSauce) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SweetandSourSauce) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (SweetandSourSauce) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SweetandSourSauce) Edible() bool {
	return true
}

// EncodeItem ...
func (SweetandSourSauce) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sweetand_sour_sauce", 0
}

// Name ...
func (SweetandSourSauce) Name() string {
	return "Sweetand Sour Sauce"
}

// Texture ...
func (SweetandSourSauce) Texture() image.Image {
	return textureFromName("sauce1")
}
