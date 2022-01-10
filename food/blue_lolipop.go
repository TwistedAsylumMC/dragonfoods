package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BlueLolipop struct{}

// AlwaysConsumable ...
func (BlueLolipop) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BlueLolipop) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BlueLolipop) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BlueLolipop) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BlueLolipop) Edible() bool {
	return true
}

// EncodeItem ...
func (BlueLolipop) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blue_lolipop", 0
}

// Name ...
func (BlueLolipop) Name() string {
	return "Blue Lolipop"
}

// Texture ...
func (BlueLolipop) Texture() image.Image {
	return textureFromName("lolipop2")
}
