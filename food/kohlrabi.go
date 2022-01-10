package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Kohlrabi struct{}

// AlwaysConsumable ...
func (Kohlrabi) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Kohlrabi) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Kohlrabi) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Kohlrabi) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Kohlrabi) Edible() bool {
	return true
}

// EncodeItem ...
func (Kohlrabi) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:kohlrabi", 0
}

// Name ...
func (Kohlrabi) Name() string {
	return "Kohlrabi"
}

// Texture ...
func (Kohlrabi) Texture() image.Image {
	return textureFromName("kohlrabi_01")
}
