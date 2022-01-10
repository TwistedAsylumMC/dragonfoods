package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Papaya struct{}

// AlwaysConsumable ...
func (Papaya) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Papaya) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Papaya) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Papaya) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Papaya) Edible() bool {
	return true
}

// EncodeItem ...
func (Papaya) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:papaya", 0
}

// Name ...
func (Papaya) Name() string {
	return "Papaya"
}

// Texture ...
func (Papaya) Texture() image.Image {
	return textureFromName("papaya_01")
}
