package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RedGummyBear struct{}

// AlwaysConsumable ...
func (RedGummyBear) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RedGummyBear) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RedGummyBear) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (RedGummyBear) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RedGummyBear) Edible() bool {
	return true
}

// EncodeItem ...
func (RedGummyBear) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:red_gummy_bear", 0
}

// Name ...
func (RedGummyBear) Name() string {
	return "RedGummyBear"
}

// Texture ...
func (RedGummyBear) Texture() image.Image {
	return textureFromName("50_giantgummybear")
}
