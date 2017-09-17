package bat

import "math"

// Pages splits long range into smaller pages.
func Pages(step int, length int) []Page {
	if step <= 0 {
		panic("Step must be a positive number.")
	}
	if length < 0 {
		panic("Step must be a positive number.")
	}

	pagesNum := int(math.Ceil(float64(length) / float64(step)))
	pages := make([]Page, pagesNum)
	for i := 0; i < pagesNum; i++ {
		pages[i].From = i * step
		pages[i].To = (i + 1) * step
	}
	if pagesNum > 0 {
		pages[pagesNum-1].To = length
	}
	return pages
}

// Page represents single range
type Page struct {
	From int
	To   int
}
