package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type InstantRamen struct{}

// AlwaysConsumable ...
func (InstantRamen) AlwaysConsumable() bool {
	return false
}

// Category ...
func (InstantRamen) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (InstantRamen) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (InstantRamen) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (InstantRamen) Edible() bool {
	return true
}

// EncodeItem ...
func (InstantRamen) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:instant_ramen", 0
}

// Name ...
func (InstantRamen) Name() string {
	return "InstantRamen"
}

// Texture ...
func (InstantRamen) Texture() image.Image {
	return textureFromName("instantramen")
}
