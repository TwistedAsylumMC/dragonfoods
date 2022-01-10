package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Salsa struct{}

// AlwaysConsumable ...
func (Salsa) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Salsa) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Salsa) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (Salsa) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Salsa) Edible() bool {
	return true
}

// EncodeItem ...
func (Salsa) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:salsa", 0
}

// Name ...
func (Salsa) Name() string {
	return "Salsa"
}

// Texture ...
func (Salsa) Texture() image.Image {
	return textureFromName("32salsa")
}
