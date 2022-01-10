package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BoxOfCereal struct{}

// AlwaysConsumable ...
func (BoxOfCereal) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BoxOfCereal) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BoxOfCereal) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(12, 7.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (BoxOfCereal) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BoxOfCereal) Edible() bool {
	return true
}

// EncodeItem ...
func (BoxOfCereal) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:box_of_cereal", 0
}

// Name ...
func (BoxOfCereal) Name() string {
	return "Box Of Cereal"
}

// Texture ...
func (BoxOfCereal) Texture() image.Image {
	return textureFromName("cereal1")
}
