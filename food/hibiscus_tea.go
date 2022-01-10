package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type HibiscusTea struct{}

// AlwaysConsumable ...
func (HibiscusTea) AlwaysConsumable() bool {
	return false
}

// Category ...
func (HibiscusTea) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (HibiscusTea) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (HibiscusTea) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (HibiscusTea) Edible() bool {
	return true
}

// EncodeItem ...
func (HibiscusTea) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:hibiscus_tea", 0
}

// Name ...
func (HibiscusTea) Name() string {
	return "HibiscusTea"
}

// Texture ...
func (HibiscusTea) Texture() image.Image {
	return textureFromName("tea_02")
}
