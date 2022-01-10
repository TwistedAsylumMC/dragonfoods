package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PinkLolipop struct{}

// AlwaysConsumable ...
func (PinkLolipop) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PinkLolipop) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PinkLolipop) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (PinkLolipop) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PinkLolipop) Edible() bool {
	return true
}

// EncodeItem ...
func (PinkLolipop) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pink_lolipop", 0
}

// Name ...
func (PinkLolipop) Name() string {
	return "PinkLolipop"
}

// Texture ...
func (PinkLolipop) Texture() image.Image {
	return textureFromName("lolipop5")
}
