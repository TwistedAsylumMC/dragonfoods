package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RawSausage struct{}

// AlwaysConsumable ...
func (RawSausage) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RawSausage) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RawSausage) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(3, 1.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (RawSausage) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RawSausage) Edible() bool {
	return true
}

// EncodeItem ...
func (RawSausage) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:raw_sausage", 0
}

// Name ...
func (RawSausage) Name() string {
	return "Raw Sausage"
}

// Texture ...
func (RawSausage) Texture() image.Image {
	return textureFromName("raw_sasuage")
}
