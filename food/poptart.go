package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Poptart struct{}

// AlwaysConsumable ...
func (Poptart) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Poptart) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Poptart) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Poptart) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Poptart) Edible() bool {
	return true
}

// EncodeItem ...
func (Poptart) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:poptart", 0
}

// Name ...
func (Poptart) Name() string {
	return "Poptart"
}

// Texture ...
func (Poptart) Texture() image.Image {
	return textureFromName("poptart")
}
