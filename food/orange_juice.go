package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type OrangeJuice struct{}

// AlwaysConsumable ...
func (OrangeJuice) AlwaysConsumable() bool {
	return false
}

// Category ...
func (OrangeJuice) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (OrangeJuice) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (OrangeJuice) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (OrangeJuice) Edible() bool {
	return true
}

// EncodeItem ...
func (OrangeJuice) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:orange_juice", 0
}

// Name ...
func (OrangeJuice) Name() string {
	return "OrangeJuice"
}

// Texture ...
func (OrangeJuice) Texture() image.Image {
	return textureFromName("orange_juice")
}
