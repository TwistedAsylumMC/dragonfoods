package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type TomatoSoup struct{}

// AlwaysConsumable ...
func (TomatoSoup) AlwaysConsumable() bool {
	return false
}

// Category ...
func (TomatoSoup) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (TomatoSoup) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (TomatoSoup) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (TomatoSoup) Edible() bool {
	return true
}

// EncodeItem ...
func (TomatoSoup) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:tomato_soup", 0
}

// Name ...
func (TomatoSoup) Name() string {
	return "TomatoSoup"
}

// Texture ...
func (TomatoSoup) Texture() image.Image {
	return textureFromName("tomatosoup")
}
