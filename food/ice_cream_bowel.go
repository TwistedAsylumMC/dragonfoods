package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type IceCreamBowel struct{}

// AlwaysConsumable ...
func (IceCreamBowel) AlwaysConsumable() bool {
	return false
}

// Category ...
func (IceCreamBowel) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (IceCreamBowel) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (IceCreamBowel) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (IceCreamBowel) Edible() bool {
	return true
}

// EncodeItem ...
func (IceCreamBowel) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:ice_cream_bowel", 0
}

// Name ...
func (IceCreamBowel) Name() string {
	return "Ice Cream Bowel"
}

// Texture ...
func (IceCreamBowel) Texture() image.Image {
	return textureFromName("58_icecream_bowl")
}
