package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BrusselsSprout struct{}

// AlwaysConsumable ...
func (BrusselsSprout) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BrusselsSprout) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BrusselsSprout) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BrusselsSprout) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BrusselsSprout) Edible() bool {
	return true
}

// EncodeItem ...
func (BrusselsSprout) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:brussels_sprout", 0
}

// Name ...
func (BrusselsSprout) Name() string {
	return "BrusselsSprout"
}

// Texture ...
func (BrusselsSprout) Texture() image.Image {
	return textureFromName("brussels_sprout")
}
