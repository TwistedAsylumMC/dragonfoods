package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type RomaineLettuce struct{}

// AlwaysConsumable ...
func (RomaineLettuce) AlwaysConsumable() bool {
	return false
}

// Category ...
func (RomaineLettuce) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (RomaineLettuce) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(4, 2.4)
	return item.Stack{}
}

// ConsumeDuration ...
func (RomaineLettuce) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (RomaineLettuce) Edible() bool {
	return true
}

// EncodeItem ...
func (RomaineLettuce) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:romaine_lettuce", 0
}

// Name ...
func (RomaineLettuce) Name() string {
	return "RomaineLettuce"
}

// Texture ...
func (RomaineLettuce) Texture() image.Image {
	return textureFromName("lettuce_romaine_01")
}
