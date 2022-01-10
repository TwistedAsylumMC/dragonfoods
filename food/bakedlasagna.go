package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Bakedlasagna struct{}

// AlwaysConsumable ...
func (Bakedlasagna) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Bakedlasagna) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Bakedlasagna) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Bakedlasagna) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Bakedlasagna) Edible() bool {
	return true
}

// EncodeItem ...
func (Bakedlasagna) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bakedlasagna", 0
}

// Name ...
func (Bakedlasagna) Name() string {
	return "Bakedlasagna"
}

// Texture ...
func (Bakedlasagna) Texture() image.Image {
	return textureFromName("bakedlasagna")
}
