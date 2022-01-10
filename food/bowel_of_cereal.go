package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BowelOfCereal struct{}

// AlwaysConsumable ...
func (BowelOfCereal) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BowelOfCereal) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BowelOfCereal) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(5, 3.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (BowelOfCereal) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BowelOfCereal) Edible() bool {
	return true
}

// EncodeItem ...
func (BowelOfCereal) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bowel_of_cereal", 0
}

// Name ...
func (BowelOfCereal) Name() string {
	return "BowelOfCereal"
}

// Texture ...
func (BowelOfCereal) Texture() image.Image {
	return textureFromName("cerealbowel")
}
