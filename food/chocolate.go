package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Chocolate struct{}

// AlwaysConsumable ...
func (Chocolate) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Chocolate) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Chocolate) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Chocolate) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Chocolate) Edible() bool {
	return true
}

// EncodeItem ...
func (Chocolate) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chocolate", 0
}

// Name ...
func (Chocolate) Name() string {
	return "Chocolate"
}

// Texture ...
func (Chocolate) Texture() image.Image {
	return textureFromName("chocolate")
}
