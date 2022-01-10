package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Flan struct{}

// AlwaysConsumable ...
func (Flan) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Flan) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Flan) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (Flan) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Flan) Edible() bool {
	return true
}

// EncodeItem ...
func (Flan) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:flan", 0
}

// Name ...
func (Flan) Name() string {
	return "Flan"
}

// Texture ...
func (Flan) Texture() image.Image {
	return textureFromName("flan")
}
