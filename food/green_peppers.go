package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GreenPeppers struct{}

// AlwaysConsumable ...
func (GreenPeppers) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GreenPeppers) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GreenPeppers) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (GreenPeppers) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GreenPeppers) Edible() bool {
	return true
}

// EncodeItem ...
func (GreenPeppers) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:green_peppers", 0
}

// Name ...
func (GreenPeppers) Name() string {
	return "GreenPeppers"
}

// Texture ...
func (GreenPeppers) Texture() image.Image {
	return textureFromName("peppergreen")
}
