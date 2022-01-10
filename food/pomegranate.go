package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Pomegranate struct{}

// AlwaysConsumable ...
func (Pomegranate) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Pomegranate) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Pomegranate) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Pomegranate) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Pomegranate) Edible() bool {
	return true
}

// EncodeItem ...
func (Pomegranate) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pomegranate", 0
}

// Name ...
func (Pomegranate) Name() string {
	return "Pomegranate"
}

// Texture ...
func (Pomegranate) Texture() image.Image {
	return textureFromName("pomegranate_01")
}
