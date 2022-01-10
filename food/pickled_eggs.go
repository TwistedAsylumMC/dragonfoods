package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PickledEggs struct{}

// AlwaysConsumable ...
func (PickledEggs) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PickledEggs) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PickledEggs) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(9, 5.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (PickledEggs) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PickledEggs) Edible() bool {
	return true
}

// EncodeItem ...
func (PickledEggs) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:pickled_eggs", 0
}

// Name ...
func (PickledEggs) Name() string {
	return "Pickled Eggs"
}

// Texture ...
func (PickledEggs) Texture() image.Image {
	return textureFromName("pickledeggs")
}
