package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedOmlet struct{}

// AlwaysConsumable ...
func (CookedOmlet) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedOmlet) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedOmlet) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(13, 7.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedOmlet) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedOmlet) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedOmlet) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_omlet", 0
}

// Name ...
func (CookedOmlet) Name() string {
	return "Cooked Omlet"
}

// Texture ...
func (CookedOmlet) Texture() image.Image {
	return textureFromName("74_omlet_dish")
}
