package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BanhMi struct{}

// AlwaysConsumable ...
func (BanhMi) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BanhMi) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BanhMi) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (BanhMi) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BanhMi) Edible() bool {
	return true
}

// EncodeItem ...
func (BanhMi) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:banh_mi", 0
}

// Name ...
func (BanhMi) Name() string {
	return "BanhMi"
}

// Texture ...
func (BanhMi) Texture() image.Image {
	return textureFromName("burger5")
}
