package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GalloPinto struct{}

// AlwaysConsumable ...
func (GalloPinto) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GalloPinto) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GalloPinto) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (GalloPinto) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GalloPinto) Edible() bool {
	return true
}

// EncodeItem ...
func (GalloPinto) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:gallo_pinto", 0
}

// Name ...
func (GalloPinto) Name() string {
	return "GalloPinto"
}

// Texture ...
func (GalloPinto) Texture() image.Image {
	return textureFromName("gallopinto")
}
