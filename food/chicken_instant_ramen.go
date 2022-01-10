package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type ChickenInstantRamen struct{}

// AlwaysConsumable ...
func (ChickenInstantRamen) AlwaysConsumable() bool {
	return false
}

// Category ...
func (ChickenInstantRamen) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (ChickenInstantRamen) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (ChickenInstantRamen) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (ChickenInstantRamen) Edible() bool {
	return true
}

// EncodeItem ...
func (ChickenInstantRamen) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:chicken_instant_ramen", 0
}

// Name ...
func (ChickenInstantRamen) Name() string {
	return "ChickenInstantRamen"
}

// Texture ...
func (ChickenInstantRamen) Texture() image.Image {
	return textureFromName("instantramen3")
}
