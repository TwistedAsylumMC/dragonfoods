package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type YellowLemon struct{}

// AlwaysConsumable ...
func (YellowLemon) AlwaysConsumable() bool {
	return false
}

// Category ...
func (YellowLemon) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (YellowLemon) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (YellowLemon) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (YellowLemon) Edible() bool {
	return true
}

// EncodeItem ...
func (YellowLemon) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:yellow_lemon", 0
}

// Name ...
func (YellowLemon) Name() string {
	return "YellowLemon"
}

// Texture ...
func (YellowLemon) Texture() image.Image {
	return textureFromName("yellow_lemon")
}
