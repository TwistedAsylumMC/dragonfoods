package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedFishSteak struct{}

// AlwaysConsumable ...
func (CookedFishSteak) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedFishSteak) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedFishSteak) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(9, 5.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedFishSteak) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedFishSteak) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedFishSteak) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_fish_steak", 0
}

// Name ...
func (CookedFishSteak) Name() string {
	return "Cooked Fish Steak"
}

// Texture ...
func (CookedFishSteak) Texture() image.Image {
	return textureFromName("fishsteak2")
}
