package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CanOfRabbitStew struct{}

// AlwaysConsumable ...
func (CanOfRabbitStew) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CanOfRabbitStew) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CanOfRabbitStew) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(8, 4.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (CanOfRabbitStew) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CanOfRabbitStew) Edible() bool {
	return true
}

// EncodeItem ...
func (CanOfRabbitStew) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:can_of_rabbit_stew", 0
}

// Name ...
func (CanOfRabbitStew) Name() string {
	return "CanOfRabbitStew"
}

// Texture ...
func (CanOfRabbitStew) Texture() image.Image {
	return textureFromName("can10")
}
