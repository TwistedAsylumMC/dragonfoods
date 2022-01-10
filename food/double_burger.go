package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type DoubleBurger struct{}

// AlwaysConsumable ...
func (DoubleBurger) AlwaysConsumable() bool {
	return false
}

// Category ...
func (DoubleBurger) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (DoubleBurger) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (DoubleBurger) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (DoubleBurger) Edible() bool {
	return true
}

// EncodeItem ...
func (DoubleBurger) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:double_burger", 0
}

// Name ...
func (DoubleBurger) Name() string {
	return "DoubleBurger"
}

// Texture ...
func (DoubleBurger) Texture() image.Image {
	return textureFromName("burger22")
}
