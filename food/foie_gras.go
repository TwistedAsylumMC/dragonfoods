package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type FoieGras struct{}

// AlwaysConsumable ...
func (FoieGras) AlwaysConsumable() bool {
	return false
}

// Category ...
func (FoieGras) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (FoieGras) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (FoieGras) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (FoieGras) Edible() bool {
	return true
}

// EncodeItem ...
func (FoieGras) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:foie_gras", 0
}

// Name ...
func (FoieGras) Name() string {
	return "FoieGras"
}

// Texture ...
func (FoieGras) Texture() image.Image {
	return textureFromName("foiegras")
}
