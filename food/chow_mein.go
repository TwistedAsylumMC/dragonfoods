package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChowMein struct{}

// AlwaysConsumable ...
func (ChowMein) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChowMein) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChowMein) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChowMein) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChowMein) Edible() bool {
	return true
}

// EncodeItem ...
func (ChowMein) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chow_mein", 0
}

// Name ...
func (ChowMein) Name() string {
	return "Chow Mein"
}

// Texture ...
func (ChowMein) Texture() image.Image {
	return textureFromName("chowmein")
}
