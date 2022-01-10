package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Peas struct{}

// AlwaysConsumable ...
func (Peas) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Peas) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Peas) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Peas) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Peas) Edible() bool {
	return true
}

// EncodeItem ...
func (Peas) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:peas", 0
}

// Name ...
func (Peas) Name() string {
	return "Peas"
}

// Texture ...
func (Peas) Texture() image.Image {
	return textureFromName("peas")
}
