package persistence

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/yongwoon/echo-blog/form"
	"github.com/yongwoon/echo-blog/model"
	"github.com/yongwoon/echo-blog/test"
)

func TestPostPersistenceGetAll(t *testing.T) {
	Convey("GetAll", t, func() {
		Convey("全ての post を取得", func() {
			test.Setup()
			test.ImportFixture()
			pp := NewPostPersistence()

			_, count, err := pp.GetAll()
			So(count, ShouldEqual, 3)
			So(err, ShouldBeNil)
		})

		Reset(func() {
			test.Teardown()
		})

		Convey("post がない場合", func() {
			pp := NewPostPersistence()

			_, count, err := pp.GetAll()

			So(count, ShouldEqual, 0)
			So(err, ShouldBeNil)
		})

		Reset(func() {
			test.Teardown()
		})
	})
}

func TestPostPersistenceGetByID(t *testing.T) {
	Convey("GetByID", t, func() {
		test.Setup()
		test.ImportFixture()

		Convey("対象の post がある場合", func() {
			postID := int64(1)
			pp := NewPostPersistence()

			post, err := pp.GetByID(postID)
			Convey("対象のpost が取得できる", func() {
				So(post.ID, ShouldEqual, postID)
				So(err, ShouldBeNil)
			})
		})

		Convey("対象の post がない場合", func() {
			postID := int64(99)
			pp := NewPostPersistence()

			post, err := pp.GetByID(postID)

			Convey("取得したユーザが存在しない", func() {
				So(post, ShouldBeNil)
				So(err, ShouldBeNil)
			})
		})

		Reset(func() {
			test.Teardown()
		})
	})
}

func TestPostPersistenceCreate(t *testing.T) {
	Convey("Create", t, func() {
		test.Setup()

		var post model.Post
		createReq := form.PostCreateReq{
			Post: struct {
				Title string `json:"title" form:"string" valid:"required, type(string)"`
				Body  string `json:"body" form:"string" valid:"required, type(string)"`
			}{
				Title: "test title",
				Body:  "test body",
			},
		}

		pp := NewPostPersistence()

		Convey("正常にユーザーが生成される", func() {
			err := pp.Create(&post, &createReq)
			So(err, ShouldBeNil)
		})

		Reset(func() {
			test.Teardown()
		})
	})
}

func TestPostPersistenceUpdate(t *testing.T) {
	Convey("Update", t, func() {
		test.Setup()
		test.ImportFixture()

		updateReq := form.PostUpdateReq{
			Post: struct {
				Title string `json:"title" form:"string" valid:"required, type(string)"`
				Body  string `json:"body" form:"string" valid:"required, type(string)"`
			}{
				Title: "updated title",
				Body:  "updated body",
			},
		}

		postID := int64(1)

		pp := NewPostPersistence()
		post, _ := pp.GetByID(postID)

		err := pp.Update(post, &updateReq)

		Convey("エラーがない", func() {
			So(err, ShouldBeNil)
		})

		Convey("正常に生成される", func() {
			So(post.Title, ShouldEqual, "updated title")
			So(post.Body, ShouldEqual, "updated body")
		})

		Reset(func() {
			test.Teardown()
		})
	})
}

func TestPostPersistenceDelete(t *testing.T) {
	Convey("Delete", t, func() {
		test.Setup()
		test.ImportFixture()

		postID := int64(1)

		pp := NewPostPersistence()
		post, _ := pp.GetByID(postID)

		err := pp.Delete(post)

		Convey("正常に削除される", func() {
			So(err, ShouldBeNil)

			post, err := pp.GetByID(postID)
			So(err, ShouldBeNil)
			So(post, ShouldBeNil)
		})

		Reset(func() {
			test.Teardown()
		})
	})
}
