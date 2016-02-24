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
	"regexp"
	"strconv"
	"time"
)

func date(s string) time.Time {
	now := time.Now()

	r1 := regexp.MustCompile(`(\d{1,2})分钟前`).FindStringSubmatch(s)
	r2 := regexp.MustCompile(`今天 (\d{2}):(\d{2})`).FindStringSubmatch(s)
	r3 := regexp.MustCompile(`(\d{1,2})月(\d{1,2})日 (\d{2}):(\d{2})`).FindStringSubmatch(s)
	switch {
	case len(r1) > 0:
		min, _ := strconv.Atoi(r1[1])
		return now.Add(time.Duration(min) * time.Minute * -1)
	case len(r2) > 0:
		y, m, d := now.Date()
		h, _ := strconv.Atoi(r2[1])
		mm, _ := strconv.Atoi(r2[2])
		return time.Date(y, m, d, h, mm, 0, 0, now.Location())
	case len(r3) > 0:
		m, _ := strconv.Atoi(r3[1])
		d, _ := strconv.Atoi(r3[2])
		h, _ := strconv.Atoi(r3[3])
		mm, _ := strconv.Atoi(r3[4])
		return time.Date(now.Year(), time.Month(m), d, h, mm, 0, 0, now.Location())
	default:
		t, _ := time.ParseInLocation("2006-1-2 15:04", s, now.Location())
		return t
	}
}
