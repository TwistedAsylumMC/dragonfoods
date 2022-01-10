package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedEel struct{}

// AlwaysConsumable ...
func (CookedEel) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedEel) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedEel) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedEel) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedEel) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedEel) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_eel", 0
}

// Name ...
func (CookedEel) Name() string {
	return "CookedEel"
}

// Texture ...
func (CookedEel) Texture() image.Image {
	return textureFromName("cookedeel")
}
