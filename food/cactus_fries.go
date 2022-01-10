package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CactusFries struct{}

// AlwaysConsumable ...
func (CactusFries) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CactusFries) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CactusFries) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (CactusFries) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CactusFries) Edible() bool {
	return true
}

// EncodeItem ...
func (CactusFries) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cactus_fries", 0
}

// Name ...
func (CactusFries) Name() string {
	return "CactusFries"
}

// Texture ...
func (CactusFries) Texture() image.Image {
	return textureFromName("cactusfries")
}
