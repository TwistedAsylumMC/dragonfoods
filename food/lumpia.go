package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Lumpia struct{}

// AlwaysConsumable ...
func (Lumpia) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Lumpia) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Lumpia) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Lumpia) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Lumpia) Edible() bool {
	return true
}

// EncodeItem ...
func (Lumpia) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:lumpia", 0
}

// Name ...
func (Lumpia) Name() string {
	return "Lumpia"
}

// Texture ...
func (Lumpia) Texture() image.Image {
	return textureFromName("lumpia")
}
