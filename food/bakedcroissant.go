package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Bakedcroissant struct{}

// AlwaysConsumable ...
func (Bakedcroissant) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Bakedcroissant) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Bakedcroissant) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Bakedcroissant) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Bakedcroissant) Edible() bool {
	return true
}

// EncodeItem ...
func (Bakedcroissant) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bakedcroissant", 0
}

// Name ...
func (Bakedcroissant) Name() string {
	return "Bakedcroissant"
}

// Texture ...
func (Bakedcroissant) Texture() image.Image {
	return textureFromName("bakedcroissant")
}
