package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SushiRoll struct{}

// AlwaysConsumable ...
func (SushiRoll) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SushiRoll) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SushiRoll) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (SushiRoll) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SushiRoll) Edible() bool {
	return true
}

// EncodeItem ...
func (SushiRoll) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sushi_roll", 0
}

// Name ...
func (SushiRoll) Name() string {
	return "Sushi Roll"
}

// Texture ...
func (SushiRoll) Texture() image.Image {
	return textureFromName("sushi")
}
