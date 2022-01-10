package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BowelOfVanillaIceCream struct{}

// AlwaysConsumable ...
func (BowelOfVanillaIceCream) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BowelOfVanillaIceCream) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BowelOfVanillaIceCream) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BowelOfVanillaIceCream) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BowelOfVanillaIceCream) Edible() bool {
	return true
}

// EncodeItem ...
func (BowelOfVanillaIceCream) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bowel_of_vanilla_ice_cream", 0
}

// Name ...
func (BowelOfVanillaIceCream) Name() string {
	return "BowelOfVanillaIceCream"
}

// Texture ...
func (BowelOfVanillaIceCream) Texture() image.Image {
	return textureFromName("icreambowel2")
}
