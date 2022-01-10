package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type WhiteJello struct{}

// AlwaysConsumable ...
func (WhiteJello) AlwaysConsumable() bool {
	return false
}

// Category ...
func (WhiteJello) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (WhiteJello) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (WhiteJello) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (WhiteJello) Edible() bool {
	return true
}

// EncodeItem ...
func (WhiteJello) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:white_jello", 0
}

// Name ...
func (WhiteJello) Name() string {
	return "WhiteJello"
}

// Texture ...
func (WhiteJello) Texture() image.Image {
	return textureFromName("65jello")
}
