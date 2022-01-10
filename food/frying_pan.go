package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type FryingPan struct{}

// AlwaysConsumable ...
func (FryingPan) AlwaysConsumable() bool {
	return false
}

// Category ...
func (FryingPan) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (FryingPan) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (FryingPan) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (FryingPan) Edible() bool {
	return true
}

// EncodeItem ...
func (FryingPan) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:frying_pan", 0
}

// Name ...
func (FryingPan) Name() string {
	return "FryingPan"
}

// Texture ...
func (FryingPan) Texture() image.Image {
	return textureFromName("frying_pan")
}
