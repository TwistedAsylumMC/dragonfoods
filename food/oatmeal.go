package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Oatmeal struct{}

// AlwaysConsumable ...
func (Oatmeal) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Oatmeal) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Oatmeal) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Oatmeal) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Oatmeal) Edible() bool {
	return true
}

// EncodeItem ...
func (Oatmeal) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:oatmeal", 0
}

// Name ...
func (Oatmeal) Name() string {
	return "Oatmeal"
}

// Texture ...
func (Oatmeal) Texture() image.Image {
	return textureFromName("oatmeal5")
}
