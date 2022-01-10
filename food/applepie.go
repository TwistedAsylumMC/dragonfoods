package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Applepie struct{}

// AlwaysConsumable ...
func (Applepie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Applepie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Applepie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(3, 1.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (Applepie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Applepie) Edible() bool {
	return true
}

// EncodeItem ...
func (Applepie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:applepie", 0
}

// Name ...
func (Applepie) Name() string {
	return "Applepie"
}

// Texture ...
func (Applepie) Texture() image.Image {
	return textureFromName("applepie")
}
