package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueFruit struct{}

// AlwaysConsumable ...
func (BlueFruit) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueFruit) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueFruit) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueFruit) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueFruit) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueFruit) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_fruit", 0
}

// Name ...
func (BlueFruit) Name() string {
	return "BlueFruit"
}

// Texture ...
func (BlueFruit) Texture() image.Image {
	return textureFromName("gura_gura_fruit")
}
