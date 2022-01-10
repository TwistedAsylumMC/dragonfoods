package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChiaPudding struct{}

// AlwaysConsumable ...
func (ChiaPudding) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChiaPudding) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChiaPudding) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChiaPudding) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChiaPudding) Edible() bool {
	return true
}

// EncodeItem ...
func (ChiaPudding) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chia_pudding", 0
}

// Name ...
func (ChiaPudding) Name() string {
	return "Chia Pudding"
}

// Texture ...
func (ChiaPudding) Texture() image.Image {
	return textureFromName("pudding1")
}
