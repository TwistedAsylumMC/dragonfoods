package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Cola struct{}

// AlwaysConsumable ...
func (Cola) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Cola) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Cola) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Cola) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Cola) Edible() bool {
	return true
}

// EncodeItem ...
func (Cola) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cola", 0
}

// Name ...
func (Cola) Name() string {
	return "Cola"
}

// Texture ...
func (Cola) Texture() image.Image {
	return textureFromName("softdrink")
}
