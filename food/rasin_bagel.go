package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RasinBagel struct{}

// AlwaysConsumable ...
func (RasinBagel) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RasinBagel) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RasinBagel) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (RasinBagel) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RasinBagel) Edible() bool {
	return true
}

// EncodeItem ...
func (RasinBagel) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:rasin_bagel", 0
}

// Name ...
func (RasinBagel) Name() string {
	return "Rasin Bagel"
}

// Texture ...
func (RasinBagel) Texture() image.Image {
	return textureFromName("bagel")
}
