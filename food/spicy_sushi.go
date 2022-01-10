package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SpicySushi struct{}

// AlwaysConsumable ...
func (SpicySushi) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SpicySushi) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SpicySushi) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (SpicySushi) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SpicySushi) Edible() bool {
	return true
}

// EncodeItem ...
func (SpicySushi) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:spicy_sushi", 0
}

// Name ...
func (SpicySushi) Name() string {
	return "SpicySushi"
}

// Texture ...
func (SpicySushi) Texture() image.Image {
	return textureFromName("97_sushi")
}
