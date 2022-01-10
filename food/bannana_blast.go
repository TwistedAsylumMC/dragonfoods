package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type BannanaBlast struct{}

// AlwaysConsumable ...
func (BannanaBlast) AlwaysConsumable() bool {
	return false
}

// Category ...
func (BannanaBlast) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (BannanaBlast) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(6, 3.6)
	return item.Stack{}
}

// ConsumeDuration ...
func (BannanaBlast) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (BannanaBlast) Edible() bool {
	return true
}

// EncodeItem ...
func (BannanaBlast) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:bannana_blast", 0
}

// Name ...
func (BannanaBlast) Name() string {
	return "Bannana Blast"
}

// Texture ...
func (BannanaBlast) Texture() image.Image {
	return textureFromName("bannanablast")
}
