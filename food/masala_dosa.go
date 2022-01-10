package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MasalaDosa struct{}

// AlwaysConsumable ...
func (MasalaDosa) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MasalaDosa) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MasalaDosa) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (MasalaDosa) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MasalaDosa) Edible() bool {
	return true
}

// EncodeItem ...
func (MasalaDosa) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:masala_dosa", 0
}

// Name ...
func (MasalaDosa) Name() string {
	return "Masala Dosa"
}

// Texture ...
func (MasalaDosa) Texture() image.Image {
	return textureFromName("masaladosa")
}
