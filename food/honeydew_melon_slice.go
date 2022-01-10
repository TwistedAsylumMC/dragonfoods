package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type HoneydewMelonSlice struct{}

// AlwaysConsumable ...
func (HoneydewMelonSlice) AlwaysConsumable() bool {
	return false
}

// Category ...
func (HoneydewMelonSlice) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (HoneydewMelonSlice) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (HoneydewMelonSlice) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (HoneydewMelonSlice) Edible() bool {
	return true
}

// EncodeItem ...
func (HoneydewMelonSlice) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:honeydew_melon_slice", 0
}

// Name ...
func (HoneydewMelonSlice) Name() string {
	return "HoneydewMelonSlice"
}

// Texture ...
func (HoneydewMelonSlice) Texture() image.Image {
	return textureFromName("melonhoneydew")
}
