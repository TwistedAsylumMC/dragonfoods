package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type StickyToffeePudding struct{}

// AlwaysConsumable ...
func (StickyToffeePudding) AlwaysConsumable() bool {
	return false
}

// Category ...
func (StickyToffeePudding) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (StickyToffeePudding) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (StickyToffeePudding) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (StickyToffeePudding) Edible() bool {
	return true
}

// EncodeItem ...
func (StickyToffeePudding) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:sticky_toffee_pudding", 0
}

// Name ...
func (StickyToffeePudding) Name() string {
	return "StickyToffeePudding"
}

// Texture ...
func (StickyToffeePudding) Texture() image.Image {
	return textureFromName("76_pudding_dish")
}
