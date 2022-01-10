package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlockMistSoda struct{}

// AlwaysConsumable ...
func (BlockMistSoda) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlockMistSoda) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlockMistSoda) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlockMistSoda) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlockMistSoda) Edible() bool {
	return true
}

// EncodeItem ...
func (BlockMistSoda) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:block_mist_soda", 0
}

// Name ...
func (BlockMistSoda) Name() string {
	return "Block Mist Soda"
}

// Texture ...
func (BlockMistSoda) Texture() image.Image {
	return textureFromName("soda_can")
}
