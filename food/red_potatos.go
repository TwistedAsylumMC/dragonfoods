package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RedPotatos struct{}

// AlwaysConsumable ...
func (RedPotatos) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RedPotatos) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RedPotatos) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (RedPotatos) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RedPotatos) Edible() bool {
	return true
}

// EncodeItem ...
func (RedPotatos) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:red_potatos", 0
}

// Name ...
func (RedPotatos) Name() string {
	return "RedPotatos"
}

// Texture ...
func (RedPotatos) Texture() image.Image {
	return textureFromName("potatored")
}
