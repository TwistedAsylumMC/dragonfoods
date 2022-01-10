package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueJello struct{}

// AlwaysConsumable ...
func (BlueJello) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueJello) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueJello) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueJello) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueJello) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueJello) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_jello", 0
}

// Name ...
func (BlueJello) Name() string {
	return "Blue Jello"
}

// Texture ...
func (BlueJello) Texture() image.Image {
	return textureFromName("63jello")
}
