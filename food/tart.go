package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Tart struct{}

// AlwaysConsumable ...
func (Tart) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Tart) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Tart) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Tart) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Tart) Edible() bool {
	return true
}

// EncodeItem ...
func (Tart) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:tart", 0
}

// Name ...
func (Tart) Name() string {
	return "Tart"
}

// Texture ...
func (Tart) Texture() image.Image {
	return textureFromName("tart")
}
