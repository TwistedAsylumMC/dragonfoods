package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RedWine struct{}

// AlwaysConsumable ...
func (RedWine) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RedWine) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RedWine) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (RedWine) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RedWine) Edible() bool {
	return true
}

// EncodeItem ...
func (RedWine) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:red_wine", 0
}

// Name ...
func (RedWine) Name() string {
	return "RedWine"
}

// Texture ...
func (RedWine) Texture() image.Image {
	return textureFromName("wine")
}
