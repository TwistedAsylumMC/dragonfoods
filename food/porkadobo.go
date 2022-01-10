package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Porkadobo struct{}

// AlwaysConsumable ...
func (Porkadobo) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Porkadobo) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Porkadobo) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (Porkadobo) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Porkadobo) Edible() bool {
	return true
}

// EncodeItem ...
func (Porkadobo) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:porkadobo", 0
}

// Name ...
func (Porkadobo) Name() string {
	return "Porkadobo"
}

// Texture ...
func (Porkadobo) Texture() image.Image {
	return textureFromName("pork_adobo")
}
