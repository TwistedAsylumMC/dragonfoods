package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type StrawberryMilkShake struct{}

// AlwaysConsumable ...
func (StrawberryMilkShake) AlwaysConsumable() bool {
	return false
}

// Category ...
func (StrawberryMilkShake) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (StrawberryMilkShake) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (StrawberryMilkShake) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (StrawberryMilkShake) Edible() bool {
	return true
}

// EncodeItem ...
func (StrawberryMilkShake) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:strawberry_milk_shake", 0
}

// Name ...
func (StrawberryMilkShake) Name() string {
	return "Strawberry Milk Shake"
}

// Texture ...
func (StrawberryMilkShake) Texture() image.Image {
	return textureFromName("shake")
}
