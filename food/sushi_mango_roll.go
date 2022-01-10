package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SushiMangoRoll struct{}

// AlwaysConsumable ...
func (SushiMangoRoll) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SushiMangoRoll) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SushiMangoRoll) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (SushiMangoRoll) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SushiMangoRoll) Edible() bool {
	return true
}

// EncodeItem ...
func (SushiMangoRoll) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sushi_mango_roll", 0
}

// Name ...
func (SushiMangoRoll) Name() string {
	return "SushiMangoRoll"
}

// Texture ...
func (SushiMangoRoll) Texture() image.Image {
	return textureFromName("sushi_03_nigiri_tamago")
}
