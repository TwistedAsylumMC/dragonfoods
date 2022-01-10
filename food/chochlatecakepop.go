package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Chochlatecakepop struct{}

// AlwaysConsumable ...
func (Chochlatecakepop) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Chochlatecakepop) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Chochlatecakepop) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Chochlatecakepop) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (Chochlatecakepop) Edible() bool {
	return true
}

// EncodeItem ...
func (Chochlatecakepop) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chochlatecakepop", 0
}

// Name ...
func (Chochlatecakepop) Name() string {
	return "Chochlatecakepop"
}

// Texture ...
func (Chochlatecakepop) Texture() image.Image {
	return textureFromName("choclatecakepop")
}
