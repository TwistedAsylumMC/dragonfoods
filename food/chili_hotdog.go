package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChiliHotdog struct{}

// AlwaysConsumable ...
func (ChiliHotdog) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChiliHotdog) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChiliHotdog) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChiliHotdog) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChiliHotdog) Edible() bool {
	return true
}

// EncodeItem ...
func (ChiliHotdog) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chili_hotdog", 0
}

// Name ...
func (ChiliHotdog) Name() string {
	return "ChiliHotdog"
}

// Texture ...
func (ChiliHotdog) Texture() image.Image {
	return textureFromName("turkeyhotdog2")
}
