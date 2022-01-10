package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Eggtart struct{}

// AlwaysConsumable ...
func (Eggtart) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Eggtart) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Eggtart) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Eggtart) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Eggtart) Edible() bool {
	return true
}

// EncodeItem ...
func (Eggtart) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:eggtart", 0
}

// Name ...
func (Eggtart) Name() string {
	return "Eggtart"
}

// Texture ...
func (Eggtart) Texture() image.Image {
	return textureFromName("42_eggtart")
}
