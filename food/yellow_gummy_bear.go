package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type YellowGummyBear struct{}

// AlwaysConsumable ...
func (YellowGummyBear) AlwaysConsumable() bool {
	return false
}

// Category ...
func (YellowGummyBear) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (YellowGummyBear) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (YellowGummyBear) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (YellowGummyBear) Edible() bool {
	return true
}

// EncodeItem ...
func (YellowGummyBear) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:yellow_gummy_bear", 0
}

// Name ...
func (YellowGummyBear) Name() string {
	return "YellowGummyBear"
}

// Texture ...
func (YellowGummyBear) Texture() image.Image {
	return textureFromName("53gummybear")
}
