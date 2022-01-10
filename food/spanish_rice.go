package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SpanishRice struct{}

// AlwaysConsumable ...
func (SpanishRice) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SpanishRice) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SpanishRice) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (SpanishRice) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SpanishRice) Edible() bool {
	return true
}

// EncodeItem ...
func (SpanishRice) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:spanish_rice", 0
}

// Name ...
func (SpanishRice) Name() string {
	return "SpanishRice"
}

// Texture ...
func (SpanishRice) Texture() image.Image {
	return textureFromName("rice2")
}
