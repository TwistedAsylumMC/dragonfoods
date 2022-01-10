package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type FudgeChips struct{}

// AlwaysConsumable ...
func (FudgeChips) AlwaysConsumable() bool {
	return false
}

// Category ...
func (FudgeChips) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (FudgeChips) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (FudgeChips) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (FudgeChips) Edible() bool {
	return true
}

// EncodeItem ...
func (FudgeChips) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:fudge_chips", 0
}

// Name ...
func (FudgeChips) Name() string {
	return "FudgeChips"
}

// Texture ...
func (FudgeChips) Texture() image.Image {
	return textureFromName("snack2")
}
