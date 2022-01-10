package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type YellowLolipop struct{}

// AlwaysConsumable ...
func (YellowLolipop) AlwaysConsumable() bool {
	return false
}

// Category ...
func (YellowLolipop) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (YellowLolipop) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (YellowLolipop) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (YellowLolipop) Edible() bool {
	return true
}

// EncodeItem ...
func (YellowLolipop) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:yellow_lolipop", 0
}

// Name ...
func (YellowLolipop) Name() string {
	return "Yellow Lolipop"
}

// Texture ...
func (YellowLolipop) Texture() image.Image {
	return textureFromName("lolipop3")
}
