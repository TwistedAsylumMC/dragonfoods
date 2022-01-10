package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BirthdayCakePop struct{}

// AlwaysConsumable ...
func (BirthdayCakePop) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BirthdayCakePop) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BirthdayCakePop) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (BirthdayCakePop) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BirthdayCakePop) Edible() bool {
	return true
}

// EncodeItem ...
func (BirthdayCakePop) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:birthday_cake_pop", 0
}

// Name ...
func (BirthdayCakePop) Name() string {
	return "Birthday Cake Pop"
}

// Texture ...
func (BirthdayCakePop) Texture() image.Image {
	return textureFromName("birthdaycakepop")
}
