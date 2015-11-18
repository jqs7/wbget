//Copyright © 2015 Jqs7

//This program is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
//along with this program.  If not, see <http://www.gnu.org/licenses/>.

package getwb

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Base URL of weibo page to fetch
const WB_BASE = "http://service.weibo.com/widget/widget_blog.php?uid=%s"

// User is the struct of weibo user
type User struct {
	Name  string //the name of weibo user
	Posts []Post //the posts of weibo user
}

// Post is the struct of each post of the weibo user
type Post struct {
	// the repost text of post
	Repost     string
	RepostFrom string
	// the link of post
	Link string
	Text string
	// the thumbnail of the post's image
	Thumbnail string
	Time      time.Time
}

// Get the latest weibo information by the uid of user
func Get(uid string) (user User, err error) {
	doc, err := goquery.NewDocument(fmt.Sprintf(WB_BASE, uid))
	if err != nil {
		return User{}, err
	}
	doc.Find("#widget_wapper.mib_wgt.wdAuto div.wgtBox").
		Each(func(i int, s *goquery.Selection) {
		s.Find("div.wgtTop div.userInfo div.userNm.txt_b").
			Each(func(i int, s *goquery.Selection) {
			user.Name = s.Text()
		})
		s.Find(`div.wgtMain div#widget_content_wapper.wgtContent
		div#content_all.wgtList div.wgtCell div.wgtCell_con`).
			Each(func(i int, s *goquery.Selection) {
			p := Post{}
			s.Find("div.wgtCell_txtBot span.wgtCell_tm a.link_d").
				Each(func(i int, s *goquery.Selection) {
				p.Time = date(s.Text())
				p.Link, _ = s.Attr("href")
			})
			s.Find("p.wgtCell_txt").
				Each(func(i int, s *goquery.Selection) {
				s.Find("a img.wgt_img").Each(func(i int, s *goquery.Selection) {
					p.Thumbnail, _ = s.Attr("src")
				})
				tmp := strings.Split(s.Text(), "\n")
				out := []string{}
				for _, v := range tmp {
					v = strings.TrimSpace(v)
					if v != "" {
						out = append(out, v)
					}
				}
				if len(out) > 1 {
					reg := regexp.MustCompile(`转发了(.*)的微博：`).FindStringSubmatch(out[0])
					p.RepostFrom = reg[1]
					p.Repost = strings.TrimPrefix(out[1], "转发理由：")
					p.Text = strings.TrimPrefix(out[0], reg[0])
				} else {
					p.Text = out[0]
				}
			})
			user.Posts = append(user.Posts, p)
		})
	})
	return user, nil
}
