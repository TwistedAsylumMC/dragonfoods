package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BagOfBrownSugar struct{}

// AlwaysConsumable ...
func (BagOfBrownSugar) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BagOfBrownSugar) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BagOfBrownSugar) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (BagOfBrownSugar) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BagOfBrownSugar) Edible() bool {
	return true
}

// EncodeItem ...
func (BagOfBrownSugar) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bag_of_brown_sugar", 0
}

// Name ...
func (BagOfBrownSugar) Name() string {
	return "Bag Of Brown Sugar"
}

// Texture ...
func (BagOfBrownSugar) Texture() image.Image {
	return textureFromName("sugar3")
}
