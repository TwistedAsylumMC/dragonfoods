package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BagOfSuger struct{}

// AlwaysConsumable ...
func (BagOfSuger) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BagOfSuger) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BagOfSuger) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(10, 6.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (BagOfSuger) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BagOfSuger) Edible() bool {
	return true
}

// EncodeItem ...
func (BagOfSuger) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bag_of_suger", 0
}

// Name ...
func (BagOfSuger) Name() string {
	return "Bag Of Suger"
}

// Texture ...
func (BagOfSuger) Texture() image.Image {
	return textureFromName("sugar2")
}
