package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CreamCheeseFrosting struct{}

// AlwaysConsumable ...
func (CreamCheeseFrosting) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CreamCheeseFrosting) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CreamCheeseFrosting) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (CreamCheeseFrosting) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CreamCheeseFrosting) Edible() bool {
	return true
}

// EncodeItem ...
func (CreamCheeseFrosting) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cream_cheese_frosting", 0
}

// Name ...
func (CreamCheeseFrosting) Name() string {
	return "CreamCheeseFrosting"
}

// Texture ...
func (CreamCheeseFrosting) Texture() image.Image {
	return textureFromName("pie3")
}
