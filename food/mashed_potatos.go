package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type MashedPotatos struct{}

// AlwaysConsumable ...
func (MashedPotatos) AlwaysConsumable() bool {
	return false
}

// Category ...
func (MashedPotatos) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (MashedPotatos) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (MashedPotatos) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (MashedPotatos) Edible() bool {
	return true
}

// EncodeItem ...
func (MashedPotatos) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:mashed_potatos", 0
}

// Name ...
func (MashedPotatos) Name() string {
	return "MashedPotatos"
}

// Texture ...
func (MashedPotatos) Texture() image.Image {
	return textureFromName("oatmeal")
}
