package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GreenGummyBear struct{}

// AlwaysConsumable ...
func (GreenGummyBear) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GreenGummyBear) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GreenGummyBear) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (GreenGummyBear) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GreenGummyBear) Edible() bool {
	return true
}

// EncodeItem ...
func (GreenGummyBear) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:green_gummy_bear", 0
}

// Name ...
func (GreenGummyBear) Name() string {
	return "GreenGummyBear"
}

// Texture ...
func (GreenGummyBear) Texture() image.Image {
	return textureFromName("51gummybear")
}
