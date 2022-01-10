package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type HerbalTea struct{}

// AlwaysConsumable ...
func (HerbalTea) AlwaysConsumable() bool {
	return false
}

// Category ...
func (HerbalTea) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (HerbalTea) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (HerbalTea) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (HerbalTea) Edible() bool {
	return true
}

// EncodeItem ...
func (HerbalTea) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:herbal_tea", 0
}

// Name ...
func (HerbalTea) Name() string {
	return "HerbalTea"
}

// Texture ...
func (HerbalTea) Texture() image.Image {
	return textureFromName("tea_05")
}
