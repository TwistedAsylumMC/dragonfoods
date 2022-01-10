package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type UncookedEel struct{}

// AlwaysConsumable ...
func (UncookedEel) AlwaysConsumable() bool {
	return false
}

// Category ...
func (UncookedEel) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (UncookedEel) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (UncookedEel) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (UncookedEel) Edible() bool {
	return true
}

// EncodeItem ...
func (UncookedEel) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:uncooked_eel", 0
}

// Name ...
func (UncookedEel) Name() string {
	return "Uncooked Eel"
}

// Texture ...
func (UncookedEel) Texture() image.Image {
	return textureFromName("uncookedeel")
}
