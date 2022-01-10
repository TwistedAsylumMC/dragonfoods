package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Pretzel struct{}

// AlwaysConsumable ...
func (Pretzel) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Pretzel) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Pretzel) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Pretzel) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Pretzel) Edible() bool {
	return true
}

// EncodeItem ...
func (Pretzel) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pretzel", 0
}

// Name ...
func (Pretzel) Name() string {
	return "Pretzel"
}

// Texture ...
func (Pretzel) Texture() image.Image {
	return textureFromName("pretzel")
}
