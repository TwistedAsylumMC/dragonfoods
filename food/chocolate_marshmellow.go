package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChocolateMarshmellow struct{}

// AlwaysConsumable ...
func (ChocolateMarshmellow) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChocolateMarshmellow) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChocolateMarshmellow) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChocolateMarshmellow) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChocolateMarshmellow) Edible() bool {
	return true
}

// EncodeItem ...
func (ChocolateMarshmellow) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chocolate_marshmellow", 0
}

// Name ...
func (ChocolateMarshmellow) Name() string {
	return "ChocolateMarshmellow"
}

// Texture ...
func (ChocolateMarshmellow) Texture() image.Image {
	return textureFromName("marshmellow2")
}
