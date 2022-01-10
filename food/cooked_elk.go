package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedElk struct{}

// AlwaysConsumable ...
func (CookedElk) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedElk) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedElk) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedElk) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedElk) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedElk) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_elk", 0
}

// Name ...
func (CookedElk) Name() string {
	return "CookedElk"
}

// Texture ...
func (CookedElk) Texture() image.Image {
	return textureFromName("cookedelk")
}
