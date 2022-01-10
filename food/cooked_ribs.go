package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedRibs struct{}

// AlwaysConsumable ...
func (CookedRibs) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedRibs) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedRibs) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(9, 5.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedRibs) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedRibs) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedRibs) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_ribs", 0
}

// Name ...
func (CookedRibs) Name() string {
	return "Cooked Ribs"
}

// Texture ...
func (CookedRibs) Texture() image.Image {
	return textureFromName("ribs")
}
