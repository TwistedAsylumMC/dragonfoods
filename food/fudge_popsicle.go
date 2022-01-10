package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type FudgePopsicle struct{}

// AlwaysConsumable ...
func (FudgePopsicle) AlwaysConsumable() bool {
	return false
}

// Category ...
func (FudgePopsicle) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (FudgePopsicle) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (FudgePopsicle) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (FudgePopsicle) Edible() bool {
	return true
}

// EncodeItem ...
func (FudgePopsicle) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:fudge_popsicle", 0
}

// Name ...
func (FudgePopsicle) Name() string {
	return "FudgePopsicle"
}

// Texture ...
func (FudgePopsicle) Texture() image.Image {
	return textureFromName("ice_cream_bar_02")
}
