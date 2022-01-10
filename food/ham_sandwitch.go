package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type HamSandwitch struct{}

// AlwaysConsumable ...
func (HamSandwitch) AlwaysConsumable() bool {
	return false
}

// Category ...
func (HamSandwitch) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (HamSandwitch) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (HamSandwitch) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (HamSandwitch) Edible() bool {
	return true
}

// EncodeItem ...
func (HamSandwitch) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:ham_sandwitch", 0
}

// Name ...
func (HamSandwitch) Name() string {
	return "Ham Sandwitch"
}

// Texture ...
func (HamSandwitch) Texture() image.Image {
	return textureFromName("sandwhich3")
}
