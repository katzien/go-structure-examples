package reviews

import "time"

var DefaultReviews = []Review{
	{ID: "S_1", BeerID: 1, FirstName: "Joe", LastName: "Tribiani", Score: 5, Text: "This is good but this is not pizza!", Created: time.Date(2017, time.November, 10, 12, 36, 0, 0, time.UTC)},
	{ID: "S_2", BeerID: 2, FirstName: "Chandler", LastName: "Bing", Score: 1, Text: "I would SO NOT drink this ever again.", Created: time.Date(2017, time.October, 25, 5, 55, 0, 0, time.UTC)},
	{ID: "S_3", BeerID: 1, FirstName: "Ross", LastName: "Geller", Score: 4, Text: "Drank while on a break, was pretty good!", Created: time.Date(2017, time.October, 25, 12, 3, 0, 0, time.UTC)},
	{ID: "S_4", BeerID: 2, FirstName: "Phoebe", LastName: "Buffay", Score: 2, Text: "Wasn't that great, so I gave it to my smelly cat.", Created: time.Date(2017, time.October, 21, 16, 45, 0, 0, time.UTC)},
	{ID: "S_5", BeerID: 1, FirstName: "Monica", LastName: "Geller", Score: 5, Text: "AMAZING! Like Chandler's jokes!", Created: time.Date(2017, time.October, 22, 13, 41, 0, 0, time.UTC)},
	{ID: "S_6", BeerID: 2, FirstName: "Rachel", LastName: "Green", Score: 5, Text: "So yummy, just like my beef and custard trifle.", Created: time.Date(2017, time.October, 17, 9, 12, 0, 0, time.UTC)},
}
