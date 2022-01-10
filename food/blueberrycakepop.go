package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Blueberrycakepop struct{}

// AlwaysConsumable ...
func (Blueberrycakepop) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Blueberrycakepop) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Blueberrycakepop) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (Blueberrycakepop) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Blueberrycakepop) Edible() bool {
	return true
}

// EncodeItem ...
func (Blueberrycakepop) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:blueberrycakepop", 0
}

// Name ...
func (Blueberrycakepop) Name() string {
	return "Blueberrycakepop"
}

// Texture ...
func (Blueberrycakepop) Texture() image.Image {
	return textureFromName("bluebarrycakepop")
}
