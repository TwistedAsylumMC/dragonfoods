package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Mango struct{}

// AlwaysConsumable ...
func (Mango) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Mango) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Mango) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Mango) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Mango) Edible() bool {
	return true
}

// EncodeItem ...
func (Mango) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:mango", 0
}

// Name ...
func (Mango) Name() string {
	return "Mango"
}

// Texture ...
func (Mango) Texture() image.Image {
	return textureFromName("mango_red")
}
