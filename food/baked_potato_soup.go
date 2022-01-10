package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BakedPotatoSoup struct{}

// AlwaysConsumable ...
func (BakedPotatoSoup) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BakedPotatoSoup) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BakedPotatoSoup) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BakedPotatoSoup) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BakedPotatoSoup) Edible() bool {
	return true
}

// EncodeItem ...
func (BakedPotatoSoup) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:baked_potato_soup", 0
}

// Name ...
func (BakedPotatoSoup) Name() string {
	return "BakedPotatoSoup"
}

// Texture ...
func (BakedPotatoSoup) Texture() image.Image {
	return textureFromName("bakedpotatoesoup")
}
