package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Redlolipop struct{}

// AlwaysConsumable ...
func (Redlolipop) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Redlolipop) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Redlolipop) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Redlolipop) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Redlolipop) Edible() bool {
	return true
}

// EncodeItem ...
func (Redlolipop) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:redlolipop", 0
}

// Name ...
func (Redlolipop) Name() string {
	return "Redlolipop"
}

// Texture ...
func (Redlolipop) Texture() image.Image {
	return textureFromName("lolipop")
}
