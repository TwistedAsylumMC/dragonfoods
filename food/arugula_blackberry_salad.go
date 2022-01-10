package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ArugulaBlackberrySalad struct{}

// AlwaysConsumable ...
func (ArugulaBlackberrySalad) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ArugulaBlackberrySalad) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ArugulaBlackberrySalad) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (ArugulaBlackberrySalad) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ArugulaBlackberrySalad) Edible() bool {
	return true
}

// EncodeItem ...
func (ArugulaBlackberrySalad) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:arugula_blackberry_salad", 0
}

// Name ...
func (ArugulaBlackberrySalad) Name() string {
	return "ArugulaBlackberrySalad"
}

// Texture ...
func (ArugulaBlackberrySalad) Texture() image.Image {
	return textureFromName("arugulablackberrysalad")
}
