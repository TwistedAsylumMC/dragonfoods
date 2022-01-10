package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type FriedRice struct{}

// AlwaysConsumable ...
func (FriedRice) AlwaysConsumable() bool {
	return false
}

// Category ...
func (FriedRice) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (FriedRice) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (FriedRice) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (FriedRice) Edible() bool {
	return true
}

// EncodeItem ...
func (FriedRice) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:fried_rice", 0
}

// Name ...
func (FriedRice) Name() string {
	return "Fried Rice"
}

// Texture ...
func (FriedRice) Texture() image.Image {
	return textureFromName("friedrice")
}
