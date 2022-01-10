package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type PotatoeSoup struct{}

// AlwaysConsumable ...
func (PotatoeSoup) AlwaysConsumable() bool {
	return false
}

// Category ...
func (PotatoeSoup) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (PotatoeSoup) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (PotatoeSoup) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (PotatoeSoup) Edible() bool {
	return true
}

// EncodeItem ...
func (PotatoeSoup) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:potatoe_soup", 0
}

// Name ...
func (PotatoeSoup) Name() string {
	return "Potatoe Soup"
}

// Texture ...
func (PotatoeSoup) Texture() image.Image {
	return textureFromName("potatoesoup")
}
