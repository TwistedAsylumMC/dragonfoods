package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type WhiteLolipop struct{}

// AlwaysConsumable ...
func (WhiteLolipop) AlwaysConsumable() bool {
	return false
}

// Category ...
func (WhiteLolipop) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (WhiteLolipop) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (WhiteLolipop) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (WhiteLolipop) Edible() bool {
	return true
}

// EncodeItem ...
func (WhiteLolipop) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:white_lolipop", 0
}

// Name ...
func (WhiteLolipop) Name() string {
	return "White Lolipop"
}

// Texture ...
func (WhiteLolipop) Texture() image.Image {
	return textureFromName("lolipop6")
}
