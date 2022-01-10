package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CanOfVeggieSoup struct{}

// AlwaysConsumable ...
func (CanOfVeggieSoup) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CanOfVeggieSoup) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CanOfVeggieSoup) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (CanOfVeggieSoup) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CanOfVeggieSoup) Edible() bool {
	return true
}

// EncodeItem ...
func (CanOfVeggieSoup) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:can_of_veggie_soup", 0
}

// Name ...
func (CanOfVeggieSoup) Name() string {
	return "Can Of Veggie Soup"
}

// Texture ...
func (CanOfVeggieSoup) Texture() image.Image {
	return textureFromName("can8")
}
