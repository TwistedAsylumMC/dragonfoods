package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MreBurger struct{}

// AlwaysConsumable ...
func (MreBurger) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MreBurger) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MreBurger) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (MreBurger) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MreBurger) Edible() bool {
	return true
}

// EncodeItem ...
func (MreBurger) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:mre_burger", 0
}

// Name ...
func (MreBurger) Name() string {
	return "Mre Burger"
}

// Texture ...
func (MreBurger) Texture() image.Image {
	return textureFromName("mre4")
}
