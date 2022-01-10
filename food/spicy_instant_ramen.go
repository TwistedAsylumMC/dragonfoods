package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type SpicyInstantRamen struct{}

// AlwaysConsumable ...
func (SpicyInstantRamen) AlwaysConsumable() bool {
	return false
}

// Category ...
func (SpicyInstantRamen) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (SpicyInstantRamen) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (SpicyInstantRamen) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (SpicyInstantRamen) Edible() bool {
	return true
}

// EncodeItem ...
func (SpicyInstantRamen) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:spicy_instant_ramen", 0
}

// Name ...
func (SpicyInstantRamen) Name() string {
	return "Spicy Instant Ramen"
}

// Texture ...
func (SpicyInstantRamen) Texture() image.Image {
	return textureFromName("instantramen2")
}
