package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PinkGummyBear struct{}

// AlwaysConsumable ...
func (PinkGummyBear) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PinkGummyBear) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PinkGummyBear) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (PinkGummyBear) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PinkGummyBear) Edible() bool {
	return true
}

// EncodeItem ...
func (PinkGummyBear) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pink_gummy_bear", 0
}

// Name ...
func (PinkGummyBear) Name() string {
	return "Pink Gummy Bear"
}

// Texture ...
func (PinkGummyBear) Texture() image.Image {
	return textureFromName("54bear")
}
