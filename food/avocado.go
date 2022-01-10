package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Avocado struct{}

// AlwaysConsumable ...
func (Avocado) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Avocado) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Avocado) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Avocado) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Avocado) Edible() bool {
	return true
}

// EncodeItem ...
func (Avocado) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:avocado", 0
}

// Name ...
func (Avocado) Name() string {
	return "Avocado"
}

// Texture ...
func (Avocado) Texture() image.Image {
	return textureFromName("avocado")
}
