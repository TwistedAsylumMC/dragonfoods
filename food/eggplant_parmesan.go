package food

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/category"
	"github.com/df-mc/dragonfly/server/world"
	"image"
	"time"
)

type EggplantParmesan struct{}

// AlwaysConsumable ...
func (EggplantParmesan) AlwaysConsumable() bool {
	return false
}

// Category ...
func (EggplantParmesan) Category() category.Category {
	return category.Nature()
}

// Consume ...
func (EggplantParmesan) Consume(_ *world.World, c item.Consumer) item.Stack {
	c.Saturate(7, 4.2)
	return item.Stack{}
}

// ConsumeDuration ...
func (EggplantParmesan) ConsumeDuration() time.Duration {
	return consumeDuration(32)
}

// Edible ...
func (EggplantParmesan) Edible() bool {
	return true
}

// EncodeItem ...
func (EggplantParmesan) EncodeItem() (name string, meta int16) {
	return "dp_foods_plus:eggplant_parmesan", 0
}

// Name ...
func (EggplantParmesan) Name() string {
	return "Eggplant Parmesan"
}

// Texture ...
func (EggplantParmesan) Texture() image.Image {
	return textureFromName("eggplantpar")
}
