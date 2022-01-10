package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type VeggieTaco struct{}

// AlwaysConsumable ...
func (VeggieTaco) AlwaysConsumable() bool {
	return false
}

// Category ...
func (VeggieTaco) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (VeggieTaco) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (VeggieTaco) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (VeggieTaco) Edible() bool {
	return true
}

// EncodeItem ...
func (VeggieTaco) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:veggie_taco", 0
}

// Name ...
func (VeggieTaco) Name() string {
	return "Veggie Taco"
}

// Texture ...
func (VeggieTaco) Texture() image.Image {
	return textureFromName("taco3")
}
