package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BbqSauce struct{}

// AlwaysConsumable ...
func (BbqSauce) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BbqSauce) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BbqSauce) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(2, 1.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BbqSauce) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BbqSauce) Edible() bool {
	return true
}

// EncodeItem ...
func (BbqSauce) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bbq_sauce", 0
}

// Name ...
func (BbqSauce) Name() string {
	return "Bbq Sauce"
}

// Texture ...
func (BbqSauce) Texture() image.Image {
	return textureFromName("bbqsauce")
}
