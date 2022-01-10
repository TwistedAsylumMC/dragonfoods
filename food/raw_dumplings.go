package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RawDumplings struct{}

// AlwaysConsumable ...
func (RawDumplings) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RawDumplings) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RawDumplings) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (RawDumplings) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RawDumplings) Edible() bool {
	return true
}

// EncodeItem ...
func (RawDumplings) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:raw_dumplings", 0
}

// Name ...
func (RawDumplings) Name() string {
	return "RawDumplings"
}

// Texture ...
func (RawDumplings) Texture() image.Image {
	return textureFromName("36_dumplings")
}
