package design_browser_history

type BrowserHistory struct {
	sites []string
	pos   int
	last  int
}

// [null,null,null,null,"facebook.com","google.com","facebook.com",null,"linkedin.com","google.com","leetcode.com"]

func Constructor(homepage string) BrowserHistory {
	sites := make([]string, 1)
	sites[0] = homepage
	return BrowserHistory{
		sites: sites,
		pos:   0,
		last:  0,
	}
}

func (bh *BrowserHistory) Visit(url string) {
	bh.pos++

	if len(bh.sites) > bh.pos {
		bh.sites[bh.pos] = url
	} else {
		bh.sites = append(bh.sites, url)
	}

	bh.last = bh.pos
}

func (bh *BrowserHistory) Back(steps int) string {
	// move back
	bh.pos = max(0, bh.pos-steps)
	return bh.sites[bh.pos]
}

func (bh *BrowserHistory) Forward(steps int) string {
	// move forward
	bh.pos = min(bh.last, bh.pos+steps)
	return bh.sites[bh.pos]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
