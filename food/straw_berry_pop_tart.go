package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type StrawBerryPopTart struct{}

// AlwaysConsumable ...
func (StrawBerryPopTart) AlwaysConsumable() bool {
	return false
}

// Category ...
func (StrawBerryPopTart) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (StrawBerryPopTart) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (StrawBerryPopTart) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (StrawBerryPopTart) Edible() bool {
	return true
}

// EncodeItem ...
func (StrawBerryPopTart) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:straw_berry_pop_tart", 0
}

// Name ...
func (StrawBerryPopTart) Name() string {
	return "StrawBerryPopTart"
}

// Texture ...
func (StrawBerryPopTart) Texture() image.Image {
	return textureFromName("poptart3")
}
