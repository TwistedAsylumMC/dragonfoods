package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type KitchenKnife struct{}

// AlwaysConsumable ...
func (KitchenKnife) AlwaysConsumable() bool {
	return false
}

// Category ...
func (KitchenKnife) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (KitchenKnife) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(0, 0.0)
	return item.Stack{}
}

// ConsumeDuration ...
func (KitchenKnife) ConsumeDuration() time.Duration {
	return consumeDuration(0)
}

// Edible ...
func (KitchenKnife) Edible() bool {
	return true
}

// EncodeItem ...
func (KitchenKnife) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:kitchen_knife", 0
}

// Name ...
func (KitchenKnife) Name() string {
	return "Kitchen Knife"
}

// Texture ...
func (KitchenKnife) Texture() image.Image {
	return textureFromName("knife")
}
