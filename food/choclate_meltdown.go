package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChoclateMeltdown struct{}

// AlwaysConsumable ...
func (ChoclateMeltdown) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChoclateMeltdown) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChoclateMeltdown) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChoclateMeltdown) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChoclateMeltdown) Edible() bool {
	return true
}

// EncodeItem ...
func (ChoclateMeltdown) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:choclate_meltdown", 0
}

// Name ...
func (ChoclateMeltdown) Name() string {
	return "Choclate Meltdown"
}

// Texture ...
func (ChoclateMeltdown) Texture() image.Image {
	return textureFromName("choclatemeltdown")
}
