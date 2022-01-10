package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ButterCheese struct{}

// AlwaysConsumable ...
func (ButterCheese) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ButterCheese) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ButterCheese) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (ButterCheese) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ButterCheese) Edible() bool {
	return true
}

// EncodeItem ...
func (ButterCheese) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:butter_cheese", 0
}

// Name ...
func (ButterCheese) Name() string {
	return "ButterCheese"
}

// Texture ...
func (ButterCheese) Texture() image.Image {
	return textureFromName("cheese2")
}
