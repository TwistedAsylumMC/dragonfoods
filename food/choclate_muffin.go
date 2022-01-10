package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChoclateMuffin struct{}

// AlwaysConsumable ...
func (ChoclateMuffin) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChoclateMuffin) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChoclateMuffin) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChoclateMuffin) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChoclateMuffin) Edible() bool {
	return true
}

// EncodeItem ...
func (ChoclateMuffin) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:choclate_muffin", 0
}

// Name ...
func (ChoclateMuffin) Name() string {
	return "ChoclateMuffin"
}

// Texture ...
func (ChoclateMuffin) Texture() image.Image {
	return textureFromName("choclate_muffin")
}
