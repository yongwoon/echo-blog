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

func TestPostIndexCaseGeneral(t *testing.T) {
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
