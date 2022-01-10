package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Whiskey struct{}

// AlwaysConsumable ...
func (Whiskey) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Whiskey) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Whiskey) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Whiskey) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Whiskey) Edible() bool {
	return true
}

// EncodeItem ...
func (Whiskey) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:whiskey", 0
}

// Name ...
func (Whiskey) Name() string {
	return "Whiskey"
}

// Texture ...
func (Whiskey) Texture() image.Image {
	return textureFromName("whiskey2")
}
