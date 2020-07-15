package form

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPostNewPostForm(t *testing.T) {
	Convey("POST 登録form", t, func() {
		e := echo.New()

		Convey("パラメータが揃えた場合", func() {
			var reqJSON = `{"post":{"title":"test-title", "body":"test-body"}}`

			req := httptest.NewRequest(http.MethodPost, "/api/v1/posts", strings.NewReader(reqJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/posts")

			_, err := NewPost(c)
			Convey("成功", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("Title 未入力", func() {
			var reqJSON = `{"post":{"title":"", "body":""}}`

			req := httptest.NewRequest(http.MethodPost, "/api/v1/posts", strings.NewReader(reqJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/posts")

			_, err := NewPost(c)
			Convey("エラー発生", func() {
				So(err.Error(), ShouldEqual, "title required")
			})
		})

		Convey("Body 未入力", func() {
			var reqJSON = `{"post":{"title":"test-title", "body":""}}`

			req := httptest.NewRequest(http.MethodPost, "/api/v1/posts", strings.NewReader(reqJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/posts")

			_, err := NewPost(c)
			Convey("エラー発生", func() {
				So(err.Error(), ShouldEqual, "body required")
			})
		})
	})
}

func TestPostUpdatePostForm(t *testing.T) {
	Convey("POST 登録form", t, func() {
		e := echo.New()

		Convey("パラメータが揃えた場合", func() {
			var reqJSON = `{"post":{"title":"test-title", "body":"test-body"}}`

			req := httptest.NewRequest(http.MethodPatch, "/api/v1/posts/:id", strings.NewReader(reqJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/posts/:id")
			c.SetParamNames("id")
			c.SetParamValues("1")

			_, err := NewPost(c)
			Convey("成功", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("Title 未入力", func() {
			var reqJSON = `{"post":{"title":"", "body":""}}`

			req := httptest.NewRequest(http.MethodPatch, "/api/v1/posts/:id", strings.NewReader(reqJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/posts/:id")
			c.SetParamNames("id")
			c.SetParamValues("1")

			_, err := NewPost(c)
			Convey("エラー発生", func() {
				So(err.Error(), ShouldEqual, "title required")
			})
		})

		Convey("Body 未入力", func() {
			var reqJSON = `{"post":{"title":"test-title", "body":""}}`

			req := httptest.NewRequest(http.MethodPatch, "/api/v1/posts/:id", strings.NewReader(reqJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/posts/:id")
			c.SetParamNames("id")
			c.SetParamValues("1")

			_, err := NewPost(c)
			Convey("エラー発生", func() {
				So(err.Error(), ShouldEqual, "body required")
			})
		})
	})
}
