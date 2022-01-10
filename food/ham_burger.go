package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type HamBurger struct{}

// AlwaysConsumable ...
func (HamBurger) AlwaysConsumable() bool {
	return false
}

// Category ...
func (HamBurger) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (HamBurger) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (HamBurger) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (HamBurger) Edible() bool {
	return true
}

// EncodeItem ...
func (HamBurger) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:ham_burger", 0
}

// Name ...
func (HamBurger) Name() string {
	return "HamBurger"
}

// Texture ...
func (HamBurger) Texture() image.Image {
	return textureFromName("burger44")
}
