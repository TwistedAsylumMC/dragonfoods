package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueCheese struct{}

// AlwaysConsumable ...
func (BlueCheese) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueCheese) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueCheese) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueCheese) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueCheese) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueCheese) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_cheese", 0
}

// Name ...
func (BlueCheese) Name() string {
	return "BlueCheese"
}

// Texture ...
func (BlueCheese) Texture() image.Image {
	return textureFromName("cheese4")
}
