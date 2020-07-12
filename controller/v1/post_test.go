package v1

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestIntegerStuff(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}

// import (
// 	"net/http"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// 	. "github.com/smartystreets/goconvey/convey"
// )

// func TestPostIndex(t *testing.T) {
// 	Convey("Post 一覧が取得できる", t, func() {
// 		// db := db.DbManager()
// 		// fixtures.Import()

// 		// posts := &model.Post{}

// 		c, rec := createContext(echo.GET, "api/v1/posts")
// 		err := PostController().Index(c)
// 		So(err, ShouldBeNil)
// 		So(err, ShouldBeNil)

// 		Convey("返り値が正しい", func() {
// 			fmt.
// 				So(rec.Code, ShouldEqual, http.StatusOK)
// 		})
// 		Reset(func() {
// 			reset()
// 		})
// 	})
// }
