package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MinersCurseSoda struct{}

// AlwaysConsumable ...
func (MinersCurseSoda) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MinersCurseSoda) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MinersCurseSoda) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (MinersCurseSoda) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MinersCurseSoda) Edible() bool {
	return true
}

// EncodeItem ...
func (MinersCurseSoda) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:miners_curse_soda", 0
}

// Name ...
func (MinersCurseSoda) Name() string {
	return "MinersCurseSoda"
}

// Texture ...
func (MinersCurseSoda) Texture() image.Image {
	return textureFromName("soft_drink_red")
}
