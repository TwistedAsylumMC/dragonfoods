package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedGroundBeef struct{}

// AlwaysConsumable ...
func (CookedGroundBeef) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedGroundBeef) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedGroundBeef) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedGroundBeef) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedGroundBeef) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedGroundBeef) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_ground_beef", 0
}

// Name ...
func (CookedGroundBeef) Name() string {
	return "Cooked Ground Beef"
}

// Texture ...
func (CookedGroundBeef) Texture() image.Image {
	return textureFromName("cookedgroundbeef")
}
