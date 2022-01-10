package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type FriedChicken struct{}

// AlwaysConsumable ...
func (FriedChicken) AlwaysConsumable() bool {
	return false
}

// Category ...
func (FriedChicken) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (FriedChicken) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (FriedChicken) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (FriedChicken) Edible() bool {
	return true
}

// EncodeItem ...
func (FriedChicken) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:fried_chicken", 0
}

// Name ...
func (FriedChicken) Name() string {
	return "Fried Chicken"
}

// Texture ...
func (FriedChicken) Texture() image.Image {
	return textureFromName("friedchicken")
}
