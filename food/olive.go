package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Olive struct{}

// AlwaysConsumable ...
func (Olive) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Olive) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Olive) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Olive) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Olive) Edible() bool {
	return true
}

// EncodeItem ...
func (Olive) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:olive", 0
}

// Name ...
func (Olive) Name() string {
	return "Olive"
}

// Texture ...
func (Olive) Texture() image.Image {
	return textureFromName("olive")
}
