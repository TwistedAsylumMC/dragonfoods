package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Daikon struct{}

// AlwaysConsumable ...
func (Daikon) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Daikon) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Daikon) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Daikon) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Daikon) Edible() bool {
	return true
}

// EncodeItem ...
func (Daikon) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:daikon", 0
}

// Name ...
func (Daikon) Name() string {
	return "Daikon"
}

// Texture ...
func (Daikon) Texture() image.Image {
	return textureFromName("daikon")
}
