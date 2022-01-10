package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type UncookedElk struct{}

// AlwaysConsumable ...
func (UncookedElk) AlwaysConsumable() bool {
	return false
}

// Category ...
func (UncookedElk) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (UncookedElk) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (UncookedElk) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (UncookedElk) Edible() bool {
	return true
}

// EncodeItem ...
func (UncookedElk) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:uncooked_elk", 0
}

// Name ...
func (UncookedElk) Name() string {
	return "Uncooked Elk"
}

// Texture ...
func (UncookedElk) Texture() image.Image {
	return textureFromName("uncookedelk")
}
