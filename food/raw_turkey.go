package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RawTurkey struct{}

// AlwaysConsumable ...
func (RawTurkey) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RawTurkey) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RawTurkey) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (RawTurkey) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RawTurkey) Edible() bool {
	return true
}

// EncodeItem ...
func (RawTurkey) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:raw_turkey", 0
}

// Name ...
func (RawTurkey) Name() string {
	return "Raw Turkey"
}

// Texture ...
func (RawTurkey) Texture() image.Image {
	return textureFromName("turkeyraw")
}
