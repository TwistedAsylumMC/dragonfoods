package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type CookedDumplings struct{}

// AlwaysConsumable ...
func (CookedDumplings) AlwaysConsumable() bool {
	return false
}

// Category ...
func (CookedDumplings) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (CookedDumplings) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(13, 7.8)
	return item.Stack{}
}

// ConsumeDuration ...
func (CookedDumplings) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (CookedDumplings) Edible() bool {
	return true
}

// EncodeItem ...
func (CookedDumplings) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:cooked_dumplings", 0
}

// Name ...
func (CookedDumplings) Name() string {
	return "Cooked Dumplings"
}

// Texture ...
func (CookedDumplings) Texture() image.Image {
	return textureFromName("37_dumplings_dish")
}
