package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Pear struct{}

// AlwaysConsumable ...
func (Pear) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Pear) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Pear) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Pear) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Pear) Edible() bool {
	return true
}

// EncodeItem ...
func (Pear) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pear", 0
}

// Name ...
func (Pear) Name() string {
	return "Pear"
}

// Texture ...
func (Pear) Texture() image.Image {
	return textureFromName("pear_item")
}
