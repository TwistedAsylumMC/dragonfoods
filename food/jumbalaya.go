package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Jumbalaya struct{}

// AlwaysConsumable ...
func (Jumbalaya) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Jumbalaya) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Jumbalaya) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(14, 8.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Jumbalaya) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Jumbalaya) Edible() bool {
	return true
}

// EncodeItem ...
func (Jumbalaya) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:jumbalaya", 0
}

// Name ...
func (Jumbalaya) Name() string {
	return "Jumbalaya"
}

// Texture ...
func (Jumbalaya) Texture() image.Image {
	return textureFromName("jumbalaya")
}
