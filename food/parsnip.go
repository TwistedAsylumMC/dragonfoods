package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Parsnip struct{}

// AlwaysConsumable ...
func (Parsnip) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Parsnip) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Parsnip) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Parsnip) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Parsnip) Edible() bool {
	return true
}

// EncodeItem ...
func (Parsnip) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:parsnip", 0
}

// Name ...
func (Parsnip) Name() string {
	return "Parsnip"
}

// Texture ...
func (Parsnip) Texture() image.Image {
	return textureFromName("parsnip")
}
