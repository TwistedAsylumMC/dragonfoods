package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BannanaJam struct{}

// AlwaysConsumable ...
func (BannanaJam) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BannanaJam) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BannanaJam) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BannanaJam) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BannanaJam) Edible() bool {
	return true
}

// EncodeItem ...
func (BannanaJam) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bannana_jam", 0
}

// Name ...
func (BannanaJam) Name() string {
	return "Bannana Jam"
}

// Texture ...
func (BannanaJam) Texture() image.Image {
	return textureFromName("bannajam")
}
