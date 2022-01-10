package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Pomelo struct{}

// AlwaysConsumable ...
func (Pomelo) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Pomelo) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Pomelo) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Pomelo) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Pomelo) Edible() bool {
	return true
}

// EncodeItem ...
func (Pomelo) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pomelo", 0
}

// Name ...
func (Pomelo) Name() string {
	return "Pomelo"
}

// Texture ...
func (Pomelo) Texture() image.Image {
	return textureFromName("pomelo")
}
