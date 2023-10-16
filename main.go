package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jackdanger/collectlinks"
	"net/http"
	"regexp"
	"sort"
)

func main() {
	fmt.Println("hello world")
}

func getUrlContent(url string) {
	resp, _ := http.Get("https://www.66law.cn/laws/laodonggongshang/ldht/ldhtgl/page_2.aspx")
	links := collectlinks.All(resp.Body)
	str1, _ := regexp.Compile("^/laws/[0-9]{0,}.aspx$")
	for _, link := range links {
		if str1.MatchString(link) {
			fmt.Println(link)
			detail, _ := http.Get("https://www.66law.cn" + link)
			// Load the HTML document
			doc, err := goquery.NewDocumentFromReader(detail.Body)
			if err != nil {
				fmt.Errorf("%v", err)
			}
			// Find the review items
			doc.Find(".topkey .left-box .detail-page .det-title").Each(func(i int, s *goquery.Selection) {
				// For each item found, get the title
				title := s.Find("h1").Text()
				fmt.Printf("Review %d: %s\n", i, title)
			})
		}
	}
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var index int
	var result [][]int
	for index = 0; index <= len(nums)-3; index++ {
		var left, right int
		left = index + 1
		right = len(nums) - 1
		if nums[index] > 0 {
			break
		}
		if index > 0 && nums[index+1] == nums[index] && nums[index+1] == nums[index+2] {
			continue
		}
		for {
			if left >= right {
				break
			}
			if nums[left] == nums[left-1] {
				left++
				continue
			}
			if nums[right-1] == nums[right] {
				right--
				continue
			}
			if nums[index]+nums[left]+nums[right] == 0 {
				result = append(result, []int{nums[index], nums[left], nums[right]})
			} else if nums[index]+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
