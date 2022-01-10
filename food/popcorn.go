package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Popcorn struct{}

// AlwaysConsumable ...
func (Popcorn) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Popcorn) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Popcorn) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Popcorn) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Popcorn) Edible() bool {
	return true
}

// EncodeItem ...
func (Popcorn) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:popcorn", 0
}

// Name ...
func (Popcorn) Name() string {
	return "Popcorn"
}

// Texture ...
func (Popcorn) Texture() image.Image {
	return textureFromName("83_popcorn")
}
