package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Cheese struct{}

// AlwaysConsumable ...
func (Cheese) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Cheese) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Cheese) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(1, 0.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Cheese) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Cheese) Edible() bool {
	return true
}

// EncodeItem ...
func (Cheese) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cheese", 0
}

// Name ...
func (Cheese) Name() string {
	return "Cheese"
}

// Texture ...
func (Cheese) Texture() image.Image {
	return textureFromName("cheese")
}
