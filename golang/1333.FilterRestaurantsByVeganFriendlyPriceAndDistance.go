type Restaurant struct {
	id     int
	rating int
}

type Restaurants []Restaurant

func (restaurants Restaurants) Len() int {
	return len(restaurants)
}

func (restaurants Restaurants) Less(i, j int) bool {
	if restaurants[i].rating != restaurants[j].rating {
		return restaurants[i].rating > restaurants[j].rating
	}
	return restaurants[i].id > restaurants[j].id
}

func (restaurants Restaurants) Swap(i, j int) {
	restaurants[i], restaurants[j] = restaurants[j], restaurants[i]
}

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	var rests []Restaurant
	for _, restaurant := range restaurants {
		if restaurant[2] == 0 && veganFriendly == 1 {
			continue
		}
		if restaurant[3] > maxPrice {
			continue
		}
		if restaurant[4] > maxDistance {
			continue
		}
		rests = append(rests, Restaurant{restaurant[0], restaurant[1]})
	}
	sort.Sort(Restaurants(rests))

	var ans []int
	for _, rest := range rests {
		ans = append(ans, rest.id)
	}
	return ans
}

