package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RawGroundBeef struct{}

// AlwaysConsumable ...
func (RawGroundBeef) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RawGroundBeef) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RawGroundBeef) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(1, 0.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (RawGroundBeef) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RawGroundBeef) Edible() bool {
	return true
}

// EncodeItem ...
func (RawGroundBeef) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:raw_ground_beef", 0
}

// Name ...
func (RawGroundBeef) Name() string {
	return "RawGroundBeef"
}

// Texture ...
func (RawGroundBeef) Texture() image.Image {
	return textureFromName("ground_beef")
}
