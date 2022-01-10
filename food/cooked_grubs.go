package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedGrubs struct{}

// AlwaysConsumable ...
func (CookedGrubs) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedGrubs) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedGrubs) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedGrubs) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedGrubs) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedGrubs) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_grubs", 0
}

// Name ...
func (CookedGrubs) Name() string {
	return "Cooked Grubs"
}

// Texture ...
func (CookedGrubs) Texture() image.Image {
	return textureFromName("grubs")
}
