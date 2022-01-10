package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CanOfTomatoSoup struct{}

// AlwaysConsumable ...
func (CanOfTomatoSoup) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CanOfTomatoSoup) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CanOfTomatoSoup) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (CanOfTomatoSoup) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CanOfTomatoSoup) Edible() bool {
	return true
}

// EncodeItem ...
func (CanOfTomatoSoup) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:can_of_tomato_soup", 0
}

// Name ...
func (CanOfTomatoSoup) Name() string {
	return "CanOfTomatoSoup"
}

// Texture ...
func (CanOfTomatoSoup) Texture() image.Image {
	return textureFromName("can7")
}
