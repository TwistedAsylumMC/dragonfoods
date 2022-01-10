package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Bakedeclair struct{}

// AlwaysConsumable ...
func (Bakedeclair) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Bakedeclair) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Bakedeclair) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (Bakedeclair) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Bakedeclair) Edible() bool {
	return true
}

// EncodeItem ...
func (Bakedeclair) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bakedeclair", 0
}

// Name ...
func (Bakedeclair) Name() string {
	return "Bakedeclair"
}

// Texture ...
func (Bakedeclair) Texture() image.Image {
	return textureFromName("cute_bread_uwu")
}
