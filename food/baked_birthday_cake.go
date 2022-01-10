package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BakedBirthdayCake struct{}

// AlwaysConsumable ...
func (BakedBirthdayCake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BakedBirthdayCake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BakedBirthdayCake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (BakedBirthdayCake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BakedBirthdayCake) Edible() bool {
	return true
}

// EncodeItem ...
func (BakedBirthdayCake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:baked_birthday_cake", 0
}

// Name ...
func (BakedBirthdayCake) Name() string {
	return "BakedBirthdayCake"
}

// Texture ...
func (BakedBirthdayCake) Texture() image.Image {
	return textureFromName("birthdaycake")
}
