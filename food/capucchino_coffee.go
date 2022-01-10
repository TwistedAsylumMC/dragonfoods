package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CapucchinoCoffee struct{}

// AlwaysConsumable ...
func (CapucchinoCoffee) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CapucchinoCoffee) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CapucchinoCoffee) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (CapucchinoCoffee) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CapucchinoCoffee) Edible() bool {
	return true
}

// EncodeItem ...
func (CapucchinoCoffee) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:capucchino_coffee", 0
}

// Name ...
func (CapucchinoCoffee) Name() string {
	return "Capucchino Coffee"
}

// Texture ...
func (CapucchinoCoffee) Texture() image.Image {
	return textureFromName("capucchinocoffe")
}
