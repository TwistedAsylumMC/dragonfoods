package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BakedBrownie struct{}

// AlwaysConsumable ...
func (BakedBrownie) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BakedBrownie) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BakedBrownie) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (BakedBrownie) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BakedBrownie) Edible() bool {
	return true
}

// EncodeItem ...
func (BakedBrownie) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:baked_brownie", 0
}

// Name ...
func (BakedBrownie) Name() string {
	return "BakedBrownie"
}

// Texture ...
func (BakedBrownie) Texture() image.Image {
	return textureFromName("bakedbrownie")
}
