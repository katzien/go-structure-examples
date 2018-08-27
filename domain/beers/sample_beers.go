package beers

import "time"

var DefaultBeers = []Beer{
	{
		ID:      1,
		Name:    "Pliny the Elder",
		Brewery: "Russian River Brewing Company",
		Abv:     8,
		ShortDesc: "Pliny the Elder is brewed with Amarillo, " +
			"Centennial, CTZ, and Simcoe hops. It is well-balanced with " +
			"malt, hops, and alcohol, slightly bitter with a fresh hop " +
			"aroma of floral, citrus, and pine.",
		Created: time.Date(2017, time.October, 24, 22, 6, 0, 0, time.UTC),
	},
	{
		ID:      2,
		Name:    "Oatmeal Stout",
		Brewery: "Samuel Smith",
		Abv:     5,
		ShortDesc: "Brewed with well water (the original well at the " +
			"Old Brewery, sunk in 1758, is still in use, with the hard well " +
			"water being drawn from 85 feet underground); fermented in " +
			"‘stone Yorkshire squares’ to create an almost opaque, " +
			"wonderfully silky and smooth textured ale with a complex " +
			"medium dry palate and bittersweet finish.",
		Created: time.Date(2017, time.October, 24, 22, 12, 0, 0, time.UTC),
	},
	{
		ID:      3,
		Name:    "Märzen",
		Brewery: "Schlenkerla",
		Abv:     5,
		ShortDesc: "Bamberg's speciality, a dark, bottom fermented " +
			"smokebeer, brewed with Original Schlenkerla Smokemalt from " +
			"the Schlenkerla maltings and tapped according to old tradition " +
			"directly from the gravity-fed oakwood cask in the historical " +
			"brewery tavern.",
		Created: time.Date(2017, time.October, 24, 22, 17, 0, 0, time.UTC),
	},
	{
		ID:      4,
		Name:    "Duvel",
		Brewery: "Duvel Moortgat",
		Abv:     9,
		ShortDesc: "A Duvel is still seen as the reference among strong " +
			"golden ales. Its bouquet is lively and tickles the nose with an " +
			"element of citrus which even tends towards grapefruit thanks to " +
			"the use of only the highest-quality hop varieties.",
		Created: time.Date(2017, time.October, 24, 22, 24, 0, 0, time.UTC),
	},
	{
		ID:      5,
		Name:    "Negra",
		Brewery: "Modelo",
		Abv:     5,
		ShortDesc: "Brewed longer to enhance the flavors, this Munich " +
			"Dunkel-style Lager gives way to a rich flavor and remarkably " +
			"smooth taste.",
		Created: time.Date(2017, time.October, 24, 22, 27, 0, 0, time.UTC),
	},
	{
		ID:      6,
		Name:    "Guinness Draught",
		Brewery: "Guinness Ltd.",
		Abv:     4,
		ShortDesc: "Pours dark brown, almost black with solid lasting light brown head. " +
			"Aroma of bitter cocoa, light coffee and roasted malt. " +
			"Body is light sweet, medium bitter. " +
			"Body is light to medium, texture almost thin and carbonation average. " +
			"Finish is medium bitter cocoa with more pronounced roast flavor. Smooth drinker.",
		Created: time.Date(2017, time.October, 24, 22, 27, 0, 0, time.UTC),
	},
	{
		ID:      7,
		Name:    "XX Lager",
		Brewery: "Cuahutemoc Moctezuma",
		Abv:     4.2,
		ShortDesc: "A crisp, refreshing, light-bodied malt-flavored beer with a well-balanced finish. " +
			"A Lager that drinks like a Pilsner. A liquid embodiment of living life to the fullest. " +
			"A beverage made from pure spring water and the choicest hops. A beer with such good taste, it’s chosen you to drink it.",
		Created: time.Date(2017, time.October, 28, 15, 02, 0, 0, time.UTC),
	},
	{
		ID:      8,
		Name:    "Tecate",
		Brewery: "Cuahutemoc Moctezuma",
		Abv:     5,
		ShortDesc: "Very smooth, medium bodied brew. Malt sweetness is thin, and can be likened to diluted sugar water. " +
			"Touch of fructose-like sweetness. Light citric hop flavours gently prick the palate with tea-like notes that follow and fade quickly. " +
			"Finishes a bit dry with husk tannins and a pasty mouthfeel.",
		Created: time.Date(2017, time.October, 28, 15, 07, 0, 0, time.UTC),
	},
	{
		ID:      9,
		Name:    "Sol",
		Brewery: "Cuahutemoc Moctezuma",
		Abv:     5,
		ShortDesc: "While Corona wins the marketing wars in the U.S., Sol is the winning brand in much of Mexico, despite not being a standout in any respect. " +
			"You see the logo plastered everywhere and it’s seemingly on every restaurant and bar menu. Like Corona, it’s simple and inoffensive, " +
			"but still slightly more flavorful than your typical American macrobrew. At its best ice cold, and progressively worse as it gets warmer.",
		Created: time.Date(2017, time.October, 28, 15, 12, 0, 0, time.UTC),
	},
	{
		ID:      10,
		Name:    "Corona",
		Brewery: "Cuahutemoc Moctezuma",
		Abv:     5,
		ShortDesc: "One of the five best-selling beers in the world, but it usually tastes better in Mexico, " +
			"where the bottles don’t have so much time in transit and on shelves. (Sunlight coming through clear bottles is never a good thing for beer.) " +
			"This is the typical “drink all afternoon” beer, working well on its own or with a plate of tacos. Refreshing with a lime.",
		Created: time.Date(2017, time.October, 28, 15, 14, 0, 0, time.UTC),
	},
}
