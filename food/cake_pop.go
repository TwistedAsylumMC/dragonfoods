package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CakePop struct{}

// AlwaysConsumable ...
func (CakePop) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CakePop) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CakePop) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (CakePop) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CakePop) Edible() bool {
	return true
}

// EncodeItem ...
func (CakePop) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cake_pop", 0
}

// Name ...
func (CakePop) Name() string {
	return "CakePop"
}

// Texture ...
func (CakePop) Texture() image.Image {
	return textureFromName("cake_pop")
}
