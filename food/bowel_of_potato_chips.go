package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BowelOfPotatoChips struct{}

// AlwaysConsumable ...
func (BowelOfPotatoChips) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BowelOfPotatoChips) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BowelOfPotatoChips) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BowelOfPotatoChips) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BowelOfPotatoChips) Edible() bool {
	return true
}

// EncodeItem ...
func (BowelOfPotatoChips) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bowel_of_potato_chips", 0
}

// Name ...
func (BowelOfPotatoChips) Name() string {
	return "BowelOfPotatoChips"
}

// Texture ...
func (BowelOfPotatoChips) Texture() image.Image {
	return textureFromName("78_potatochips_bowl")
}
