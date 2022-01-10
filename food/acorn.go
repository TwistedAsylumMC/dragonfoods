package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Acorn struct{}

// AlwaysConsumable ...
func (Acorn) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Acorn) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Acorn) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Acorn) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Acorn) Edible() bool {
	return true
}

// EncodeItem ...
func (Acorn) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:acorn", 0
}

// Name ...
func (Acorn) Name() string {
	return "Acorn"
}

// Texture ...
func (Acorn) Texture() image.Image {
	return textureFromName("nut")
}
