package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Jelly struct{}

// AlwaysConsumable ...
func (Jelly) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Jelly) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Jelly) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Jelly) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (Jelly) Edible() bool {
	return true
}

// EncodeItem ...
func (Jelly) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:jelly", 0
}

// Name ...
func (Jelly) Name() string {
	return "Jelly"
}

// Texture ...
func (Jelly) Texture() image.Image {
	return textureFromName("jellyjar")
}
