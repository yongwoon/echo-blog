package v1

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yongwoon/echo-blog/persistence"
	"github.com/yongwoon/echo-blog/serializer"
	"github.com/yongwoon/echo-blog/test"
)

func TestPostIndex(t *testing.T) {
	Convey("GET api/v1/posts", t, func() {
		Convey("post が存在する場合", func() {
			e := echo.New()
			test.Setup()
			test.ImportFixture()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/posts", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/posts")

			postPersistence := persistence.NewPostPersistence()
			pc := NewPost(postPersistence)
			pc.Index(c)

			var res serializer.PostListSerializer
			test.ParseResponseBody(rec.Body, &res)

			So(rec.Code, ShouldEqual, http.StatusOK)
			So(res.PostsCount, ShouldEqual, 3)
		})

		Reset(func() {
			test.Teardown()
		})

		Convey("post が存在しない場合", func() {
			e := echo.New()
			test.Setup()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/posts", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("/api/v1/posts")

			postPersistence := persistence.NewPostPersistence()
			pc := NewPost(postPersistence)
			pc.Index(c)

			var resultJSON = `{"posts":[],"postsCount":0}`

			So(rec.Code, ShouldEqual, http.StatusOK)
			So(strings.TrimSpace(rec.Body.String()), ShouldEqual, resultJSON)
		})

		Reset(func() {
			test.Teardown()
		})
	})
}

func TestPostShow(t *testing.T) {
	Convey("GET api/v1/posts/:id", t, func() {
		e := echo.New()
		test.Setup()
		test.ImportFixture()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/posts/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/posts/:id")
		c.SetParamNames("id")

		Convey("post が存在しない場合", func() {
			c.SetParamValues("99999")

			postPersistence := persistence.NewPostPersistence()
			pc := NewPost(postPersistence)
			pc.Show(c)

			Convey("404 error", func() {
				So(rec.Code, ShouldEqual, http.StatusNotFound)
			})
		})

		Convey("post が存在する場合", func() {
			c.SetParamValues("3")

			postPersistence := persistence.NewPostPersistence()
			pc := NewPost(postPersistence)
			pc.Show(c)

			Convey("success", func() {
				So(rec.Code, ShouldEqual, http.StatusOK)
			})
		})

		Reset(func() {
			test.Teardown()
		})
	})
}

func TestPostCreate(t *testing.T) {
	Convey("POST api/v1/posts/", t, func() {
		e := echo.New()
		test.Setup()

		var reqJSON = `{"post":{"title":"test-title", "body":"test-body"}}`

		req := httptest.NewRequest(http.MethodPost, "/api/v1/posts", strings.NewReader(reqJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/posts")

		postPersistence := persistence.NewPostPersistence()
		pc := NewPost(postPersistence)
		pc.Create(c)

		var res serializer.SinglePostSerializer
		test.ParseResponseBody(rec.Body, &res)

		Convey("post が登録される", func() {
			Convey("success", func() {
				So(rec.Code, ShouldEqual, http.StatusOK)
			})

			Convey("response registerd post title", func() {
				So(res.Post.Title, ShouldEqual, "test-title")
			})
		})
	})
}

func TestPostUpdate(t *testing.T) {
	Convey("PATCH api/v1/posts/:id", t, func() {
		e := echo.New()
		test.Setup()
		test.ImportFixture()

		var reqJSON = `{"post":{"title":"updated-title", "body":"updated-body"}}`

		req := httptest.NewRequest(http.MethodPatch, "/api/v1/posts/:id", strings.NewReader(reqJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/posts")
		c.SetPath("/api/v1/posts/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		postPersistence := persistence.NewPostPersistence()
		pc := NewPost(postPersistence)
		pc.Update(c)

		var res serializer.SinglePostSerializer
		test.ParseResponseBody(rec.Body, &res)

		Convey("post が更新される", func() {
			Convey("success", func() {
				So(rec.Code, ShouldEqual, http.StatusOK)
			})

			Convey("response registerd post title", func() {
				So(res.Post.Title, ShouldEqual, "updated-title")
			})
		})

		Reset(func() {
			test.Teardown()
		})
	})
}

func TestPostDelete(t *testing.T) {
	Convey("DELETE api/v1/posts/:id", t, func() {
		e := echo.New()
		test.Setup()
		test.ImportFixture()

		req := httptest.NewRequest(http.MethodDelete, "/api/v1/posts/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/posts/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		postPersistence := persistence.NewPostPersistence()
		pc := NewPost(postPersistence)
		pc.Delete(c)

		Convey("post が削除される", func() {
			So(rec.Code, ShouldEqual, http.StatusOK)
		})

		Reset(func() {
			test.Teardown()
		})
	})
}
