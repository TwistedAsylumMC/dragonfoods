package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChoclateCake struct{}

// AlwaysConsumable ...
func (ChoclateCake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChoclateCake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChoclateCake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChoclateCake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChoclateCake) Edible() bool {
	return true
}

// EncodeItem ...
func (ChoclateCake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:choclate_cake", 0
}

// Name ...
func (ChoclateCake) Name() string {
	return "ChoclateCake"
}

// Texture ...
func (ChoclateCake) Texture() image.Image {
	return textureFromName("chocakeunbaked")
}
