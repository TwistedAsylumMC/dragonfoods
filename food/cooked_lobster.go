package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedLobster struct{}

// AlwaysConsumable ...
func (CookedLobster) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedLobster) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedLobster) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(14, 8.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedLobster) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedLobster) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedLobster) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_lobster", 0
}

// Name ...
func (CookedLobster) Name() string {
	return "CookedLobster"
}

// Texture ...
func (CookedLobster) Texture() image.Image {
	return textureFromName("lobster2")
}
