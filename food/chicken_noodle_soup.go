package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChickenNoodleSoup struct{}

// AlwaysConsumable ...
func (ChickenNoodleSoup) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChickenNoodleSoup) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChickenNoodleSoup) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChickenNoodleSoup) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChickenNoodleSoup) Edible() bool {
	return true
}

// EncodeItem ...
func (ChickenNoodleSoup) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chicken_noodle_soup", 0
}

// Name ...
func (ChickenNoodleSoup) Name() string {
	return "ChickenNoodleSoup"
}

// Texture ...
func (ChickenNoodleSoup) Texture() image.Image {
	return textureFromName("chicken_noodlesoup")
}
