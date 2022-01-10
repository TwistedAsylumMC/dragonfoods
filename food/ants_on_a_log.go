package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type AntsOnALog struct{}

// AlwaysConsumable ...
func (AntsOnALog) AlwaysConsumable() bool {
	return false
}

// Category ...
func (AntsOnALog) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (AntsOnALog) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(3, 1.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (AntsOnALog) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (AntsOnALog) Edible() bool {
	return true
}

// EncodeItem ...
func (AntsOnALog) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:ants_on_a_log", 0
}

// Name ...
func (AntsOnALog) Name() string {
	return "AntsOnALog"
}

// Texture ...
func (AntsOnALog) Texture() image.Image {
	return textureFromName("antsonalog")
}
