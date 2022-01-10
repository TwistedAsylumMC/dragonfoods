package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type GreenLolipop struct{}

// AlwaysConsumable ...
func (GreenLolipop) AlwaysConsumable() bool {
	return false
}

// Category ...
func (GreenLolipop) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (GreenLolipop) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (GreenLolipop) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (GreenLolipop) Edible() bool {
	return true
}

// EncodeItem ...
func (GreenLolipop) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:green_lolipop", 0
}

// Name ...
func (GreenLolipop) Name() string {
	return "GreenLolipop"
}

// Texture ...
func (GreenLolipop) Texture() image.Image {
	return textureFromName("lolipop4")
}
