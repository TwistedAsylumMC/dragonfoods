package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Yellowpepers struct{}

// AlwaysConsumable ...
func (Yellowpepers) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Yellowpepers) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Yellowpepers) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Yellowpepers) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Yellowpepers) Edible() bool {
	return true
}

// EncodeItem ...
func (Yellowpepers) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:yellowpepers", 0
}

// Name ...
func (Yellowpepers) Name() string {
	return "Yellowpepers"
}

// Texture ...
func (Yellowpepers) Texture() image.Image {
	return textureFromName("cat_fruit")
}
