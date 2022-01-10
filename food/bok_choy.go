package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BokChoy struct{}

// AlwaysConsumable ...
func (BokChoy) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BokChoy) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BokChoy) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BokChoy) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BokChoy) Edible() bool {
	return true
}

// EncodeItem ...
func (BokChoy) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bok_choy", 0
}

// Name ...
func (BokChoy) Name() string {
	return "Bok Choy"
}

// Texture ...
func (BokChoy) Texture() image.Image {
	return textureFromName("bok_choy_01")
}
