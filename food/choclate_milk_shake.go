package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChoclateMilkShake struct{}

// AlwaysConsumable ...
func (ChoclateMilkShake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChoclateMilkShake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChoclateMilkShake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChoclateMilkShake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChoclateMilkShake) Edible() bool {
	return true
}

// EncodeItem ...
func (ChoclateMilkShake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:choclate_milk_shake", 0
}

// Name ...
func (ChoclateMilkShake) Name() string {
	return "ChoclateMilkShake"
}

// Texture ...
func (ChoclateMilkShake) Texture() image.Image {
	return textureFromName("shake3")
}
