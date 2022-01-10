package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Marshmellow struct{}

// AlwaysConsumable ...
func (Marshmellow) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Marshmellow) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Marshmellow) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Marshmellow) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (Marshmellow) Edible() bool {
	return true
}

// EncodeItem ...
func (Marshmellow) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:marshmellow", 0
}

// Name ...
func (Marshmellow) Name() string {
	return "Marshmellow"
}

// Texture ...
func (Marshmellow) Texture() image.Image {
	return textureFromName("marshmellow")
}
