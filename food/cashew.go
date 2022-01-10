package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Cashew struct{}

// AlwaysConsumable ...
func (Cashew) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Cashew) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Cashew) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Cashew) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Cashew) Edible() bool {
	return true
}

// EncodeItem ...
func (Cashew) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cashew", 0
}

// Name ...
func (Cashew) Name() string {
	return "Cashew"
}

// Texture ...
func (Cashew) Texture() image.Image {
	return textureFromName("cashew")
}
