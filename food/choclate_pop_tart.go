package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChoclatePopTart struct{}

// AlwaysConsumable ...
func (ChoclatePopTart) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChoclatePopTart) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChoclatePopTart) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChoclatePopTart) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChoclatePopTart) Edible() bool {
	return true
}

// EncodeItem ...
func (ChoclatePopTart) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:choclate_pop_tart", 0
}

// Name ...
func (ChoclatePopTart) Name() string {
	return "ChoclatePopTart"
}

// Texture ...
func (ChoclatePopTart) Texture() image.Image {
	return textureFromName("poptart2")
}
