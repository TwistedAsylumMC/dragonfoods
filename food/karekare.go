package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type Karekare struct{}

// AlwaysConsumable ...
func (Karekare) AlwaysConsumable() bool {
	return false
}

// Category ...
func (Karekare) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (Karekare) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(15, 9.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (Karekare) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (Karekare) Edible() bool {
	return true
}

// EncodeItem ...
func (Karekare) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:karekare", 0
}

// Name ...
func (Karekare) Name() string {
	return "Karekare"
}

// Texture ...
func (Karekare) Texture() image.Image {
	return textureFromName("karekare")
}
