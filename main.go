package main

import (
	"math/rand"

	"github.com/thomas-liddle/claude-status/status"
)

var thinkingWords = []string{
	"Accomplishing", "Actioning", "Actualizing", "Architecting", "Baking",
	"Beaming", "Befuddling", "Billowing", "Blanching", "Bloviating",
	"Boogieing", "Boondoggling", "Booping", "Bootstrapping", "Brewing",
	"Bunning", "Burrowing", "Calculating", "Canoodling", "Caramelizing",
	"Cascading", "Catapulting", "Cerebrating", "Channeling", "Choreographing",
	"Churning", "Clauding", "Clobbering", "Coalescing", "Cogitating",
	"Combobulating", "Composing", "Computing", "Concocting", "Considering",
	"Contemplating", "Cooking", "Crafting", "Creating", "Crunching",
	"Crystallizing", "Cultivating", "Dashing", "Deciphering", "Deliberating",
	"Determining", "Discombobulating", "Dithering", "Doing", "Doodling",
	"Drizzling", "Ebbing", "Effecting", "Elucidating", "Embellishing",
	"Enchanting", "Envisioning", "Evaporating", "Fermenting", "Finagling",
	"Flibbertigibbeting", "Flowing", "Flummoxing", "Fluttering", "Forging",
	"Forming", "Frolicking", "Frosting", "Gallivanting", "Galloping",
	"Garnishing", "Generating", "Germinating", "Gesticulating", "Grooving",
	"Gusting", "Harmonizing", "Hatching", "Herding", "Honking",
	"Hullaballooing", "Hyperspacing", "Ideating", "Imagining", "Improvising",
	"Incubating", "Inferring", "Infusing", "Ionizing", "Jitterbugging",
	"Julienning", "Kneading", "Learning", "Leavening", "Levitating",
	"Lollygagging", "Manifesting", "Marinating", "Meandering", "Metamorphosing",
	"Misting", "Moonwalking", "Moseying", "Mulling", "Musing", "Mustering",
	"Nebulizing", "Noodling", "Nucleating", "Orbiting", "Orchestrating",
	"Osmosing", "Perambulating", "Percolating", "Perusing", "Philosophising",
	"Photosynthesizing", "Pivoting", "Pollinating", "Pondering", "Pontificating",
	"Pouncing", "Precipitating", "Prestidigitating", "Processing", "Propagating",
	"Puttering", "Puzzling", "Quantumizing", "Razzmatazzing", "Recalling",
	"Recombobulating", "Reticulating", "Roaming", "Roosting", "Ruminating",
	"Scampering", "Schlepping", "Scurrying", "Seasoning", "Shenaniganing",
	"Shimmying", "Simmering", "Skedaddling", "Sketching", "Slithering",
	"Smooshing", "Spelunking", "Sprouting", "Stewing", "Sublimating",
	"Swirling", "Swooping", "Symbioting", "Synthesizing", "Tempering",
	"Thinking", "Tinkering", "Tomfoolering", "Transfiguring", "Transmuting",
	"Undulating", "Unfurling", "Unravelling", "Vibing", "Waddling",
	"Wandering", "Warping", "Whatchamacalliting", "Whirlpooling", "Whirring",
	"Whisking", "Wibbling", "Wrangling", "Zesting", "Zigzagging",
}

func main() {
	word := thinkingWords[rand.Intn(len(thinkingWords))]
	status.SetSlackStatusForever("claude", word+"...")
}
