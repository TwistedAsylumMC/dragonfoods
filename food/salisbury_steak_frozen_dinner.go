package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SalisburySteakFrozenDinner struct{}

// AlwaysConsumable ...
func (SalisburySteakFrozenDinner) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SalisburySteakFrozenDinner) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SalisburySteakFrozenDinner) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(14, 8.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (SalisburySteakFrozenDinner) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SalisburySteakFrozenDinner) Edible() bool {
	return true
}

// EncodeItem ...
func (SalisburySteakFrozenDinner) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:salisbury_steak_frozen_dinner", 0
}

// Name ...
func (SalisburySteakFrozenDinner) Name() string {
	return "SalisburySteakFrozenDinner"
}

// Texture ...
func (SalisburySteakFrozenDinner) Texture() image.Image {
	return textureFromName("tvdinner")
}
