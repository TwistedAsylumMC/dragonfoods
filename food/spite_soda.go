package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SpiteSoda struct{}

// AlwaysConsumable ...
func (SpiteSoda) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SpiteSoda) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SpiteSoda) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (SpiteSoda) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SpiteSoda) Edible() bool {
	return true
}

// EncodeItem ...
func (SpiteSoda) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:spite_soda", 0
}

// Name ...
func (SpiteSoda) Name() string {
	return "SpiteSoda"
}

// Texture ...
func (SpiteSoda) Texture() image.Image {
	return textureFromName("soft_drink6")
}
