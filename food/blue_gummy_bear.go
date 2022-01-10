package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueGummyBear struct{}

// AlwaysConsumable ...
func (BlueGummyBear) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueGummyBear) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueGummyBear) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueGummyBear) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueGummyBear) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueGummyBear) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_gummy_bear", 0
}

// Name ...
func (BlueGummyBear) Name() string {
	return "Blue Gummy Bear"
}

// Texture ...
func (BlueGummyBear) Texture() image.Image {
	return textureFromName("52gummybear")
}
