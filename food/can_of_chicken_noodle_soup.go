package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CanOfChickenNoodleSoup struct{}

// AlwaysConsumable ...
func (CanOfChickenNoodleSoup) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CanOfChickenNoodleSoup) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CanOfChickenNoodleSoup) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (CanOfChickenNoodleSoup) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CanOfChickenNoodleSoup) Edible() bool {
	return true
}

// EncodeItem ...
func (CanOfChickenNoodleSoup) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:can_of_chicken_noodle_soup", 0
}

// Name ...
func (CanOfChickenNoodleSoup) Name() string {
	return "Can Of Chicken Noodle Soup"
}

// Texture ...
func (CanOfChickenNoodleSoup) Texture() image.Image {
	return textureFromName("can9")
}
