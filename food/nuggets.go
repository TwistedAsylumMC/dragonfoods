package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Nuggets struct{}

// AlwaysConsumable ...
func (Nuggets) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Nuggets) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Nuggets) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Nuggets) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Nuggets) Edible() bool {
	return true
}

// EncodeItem ...
func (Nuggets) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:nuggets", 0
}

// Name ...
func (Nuggets) Name() string {
	return "Nuggets"
}

// Texture ...
func (Nuggets) Texture() image.Image {
	return textureFromName("nugget")
}
