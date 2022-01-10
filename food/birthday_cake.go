package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BirthdayCake struct{}

// AlwaysConsumable ...
func (BirthdayCake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BirthdayCake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BirthdayCake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BirthdayCake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BirthdayCake) Edible() bool {
	return true
}

// EncodeItem ...
func (BirthdayCake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:birthday_cake", 0
}

// Name ...
func (BirthdayCake) Name() string {
	return "Birthday Cake"
}

// Texture ...
func (BirthdayCake) Texture() image.Image {
	return textureFromName("unbakedbirthcake")
}
