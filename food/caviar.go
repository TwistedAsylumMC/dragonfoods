package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Caviar struct{}

// AlwaysConsumable ...
func (Caviar) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Caviar) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Caviar) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Caviar) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Caviar) Edible() bool {
	return true
}

// EncodeItem ...
func (Caviar) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:caviar", 0
}

// Name ...
func (Caviar) Name() string {
	return "Caviar"
}

// Texture ...
func (Caviar) Texture() image.Image {
	return textureFromName("caviar")
}
