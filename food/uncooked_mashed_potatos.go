package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type UncookedMashedPotatos struct{}

// AlwaysConsumable ...
func (UncookedMashedPotatos) AlwaysConsumable() bool {
	return false
}

// Category ...
func (UncookedMashedPotatos) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (UncookedMashedPotatos) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (UncookedMashedPotatos) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (UncookedMashedPotatos) Edible() bool {
	return true
}

// EncodeItem ...
func (UncookedMashedPotatos) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:uncooked_mashed_potatos", 0
}

// Name ...
func (UncookedMashedPotatos) Name() string {
	return "Uncooked Mashed Potatos"
}

// Texture ...
func (UncookedMashedPotatos) Texture() image.Image {
	return textureFromName("oatmeal2")
}
