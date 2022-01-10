package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PicoDeGallo struct{}

// AlwaysConsumable ...
func (PicoDeGallo) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PicoDeGallo) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PicoDeGallo) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (PicoDeGallo) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PicoDeGallo) Edible() bool {
	return true
}

// EncodeItem ...
func (PicoDeGallo) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pico_de_gallo", 0
}

// Name ...
func (PicoDeGallo) Name() string {
	return "PicoDeGallo"
}

// Texture ...
func (PicoDeGallo) Texture() image.Image {
	return textureFromName("pico")
}
