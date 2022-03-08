package docker

import (
	"path/filepath"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetTagsByFileSuffix(t *testing.T) {
	basePath := "kajsdfno/sdafs/sdfadsf"

	Convey("testing correct path", t, func() {
		Convey("only base file", func() {
			patches := ApplyFunc(filewalk, func(_ string, callback filepath.WalkFunc) error {
				callback(basePath, nil, nil)
				return nil
			})
			defer patches.Reset()
			result := GetTagsByFileSuffix(basePath)
			So(result, ShouldResemble, []string{"latest"})
		})

		Convey("tags with one word", func() {
			tagname := "one"
			patches := ApplyFunc(filewalk, func(_ string, callback filepath.WalkFunc) error {
				callback(basePath + "." + tagname, nil, nil)
				return nil
			})
			defer patches.Reset()
			result := GetTagsByFileSuffix(basePath)
			So(result, ShouldResemble, []string{tagname})
		})
		Convey("tags with two word", func() {
			tagname := "one-two"
			patches := ApplyFunc(filewalk, func(_ string, callback filepath.WalkFunc) error {
				callback(basePath + "." + tagname, nil, nil)
				return nil
			})
			defer patches.Reset()
			result := GetTagsByFileSuffix(basePath)
			So(result, ShouldResemble, []string{tagname})
		})
		Convey("tags with multi-word", func() {
			tagname := []string{"one", "two"}
			patches := ApplyFunc(filewalk, func(_ string, callback filepath.WalkFunc) error {
				callback(basePath + "." + tagname[0], nil, nil)
				callback(basePath + "." + tagname[1], nil, nil)
				return nil
			})
			defer patches.Reset()
			result := GetTagsByFileSuffix(basePath)
			So(result, ShouldResemble, tagname)
		})
		Convey("tags with multi-word and base", func() {
			tagname := []string{"latest", "one", "two"}
			patches := ApplyFunc(filewalk, func(_ string, callback filepath.WalkFunc) error {
				callback(basePath, nil, nil)
				callback(basePath + "." + tagname[1], nil, nil)
				callback(basePath + "." + tagname[2], nil, nil)
				return nil
			})
			defer patches.Reset()
			result := GetTagsByFileSuffix(basePath)
			So(result, ShouldResemble, tagname)
		})
	})
	Convey("testing incorrect path", t, func() {
		Convey("tags without dot", func() {
			tagname := "nothing"
			patches := ApplyFunc(filewalk, func(_ string, callback filepath.WalkFunc) error {
				callback(basePath + tagname, nil, nil)
				return nil
			})
			defer patches.Reset()
			result := GetTagsByFileSuffix(basePath)
			So(result, ShouldResemble, []string{})
		})
		Convey("tags with multi wrong words", func() {
			tagname := []string{"one", "two", "latest"}
			patches := ApplyFunc(filewalk, func(_ string, callback filepath.WalkFunc) error {
				callback(basePath + "." + tagname[0], nil, nil)
				callback(basePath + "test", nil, nil)
				callback(basePath + "." + tagname[1], nil, nil)
				callback(basePath, nil, nil)
				return nil
			})
			defer patches.Reset()
			result := GetTagsByFileSuffix(basePath)
			So(result, ShouldResemble, tagname)
		})
	})
}
