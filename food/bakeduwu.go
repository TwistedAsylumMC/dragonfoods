package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Bakeduwu struct{}

// AlwaysConsumable ...
func (Bakeduwu) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Bakeduwu) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Bakeduwu) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (Bakeduwu) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Bakeduwu) Edible() bool {
	return true
}

// EncodeItem ...
func (Bakeduwu) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bakeduwu", 0
}

// Name ...
func (Bakeduwu) Name() string {
	return "Bakeduwu"
}

// Texture ...
func (Bakeduwu) Texture() image.Image {
	return textureFromName("bakedapples")
}
