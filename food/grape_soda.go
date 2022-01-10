package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GrapeSoda struct{}

// AlwaysConsumable ...
func (GrapeSoda) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GrapeSoda) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GrapeSoda) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (GrapeSoda) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GrapeSoda) Edible() bool {
	return true
}

// EncodeItem ...
func (GrapeSoda) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:grape_soda", 0
}

// Name ...
func (GrapeSoda) Name() string {
	return "GrapeSoda"
}

// Texture ...
func (GrapeSoda) Texture() image.Image {
	return textureFromName("soft_drink_blue")
}
