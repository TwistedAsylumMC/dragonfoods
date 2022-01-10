package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Tea struct{}

// AlwaysConsumable ...
func (Tea) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Tea) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Tea) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Tea) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Tea) Edible() bool {
	return true
}

// EncodeItem ...
func (Tea) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:tea", 0
}

// Name ...
func (Tea) Name() string {
	return "Tea"
}

// Texture ...
func (Tea) Texture() image.Image {
	return textureFromName("tea_01")
}
