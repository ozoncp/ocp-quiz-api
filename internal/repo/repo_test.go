package repo

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	sq "github.com/Masterminds/squirrel"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-quiz-api/internal/models"
)

var _ = Describe("Repo", func() {
	var (
		rep      Repo
		dbMock   sqlmock.Sqlmock
		mockCtrl *gomock.Controller
		ctx      context.Context
		db       *sql.DB
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		ctx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
		defer db.Close()
		if err := dbMock.ExpectationsWereMet(); err != nil {
			Expect(err).ToNot(HaveOccurred())
		}

	})
	Context("Work with DB Repo", func() {
		JustBeforeEach(func() {
			var err error
			db, dbMock, err = sqlmock.New()
			Expect(err).ToNot(HaveOccurred())
			stmtCache := sq.NewStmtCache(db)

			rep = &dbRepo{
				stmBuilder:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(stmtCache),
				tableName:   "quiz",
				userId:      "user_id",
				classroomId: "classroom_id",
				link:        "link",
			}

		})
		It("AddEntity", func() {
			newQuiz := models.Quiz{
				Id:          0,
				UserId:      1,
				ClassroomId: 2,
				Link:        "one",
			}
			expectedNewId := uint64(1)
			returnRows := sqlmock.NewRows([]string{"id"}).AddRow(expectedNewId)
			dbMock.ExpectPrepare(
				"INSERT INTO quiz \\(user_id,classroom_id,link\\) VALUES \\(\\$1,\\$2,\\$3\\) RETURNING id",
			).
				ExpectQuery().
				WithArgs(newQuiz.UserId, newQuiz.ClassroomId, newQuiz.Link).
				WillReturnRows(returnRows)

			newId, err := rep.AddEntity(ctx, newQuiz)
			Expect(err).ToNot(HaveOccurred())

			Expect(newId).To(Equal(expectedNewId))
		})

		It("AddEntities", func() {
			quiz := []models.Quiz{
				{
					UserId:      10,
					ClassroomId: 100,
					Link:        "one",
				},
				{
					UserId:      20,
					ClassroomId: 200,
					Link:        "two",
				},
				{
					UserId:      30,
					ClassroomId: 300,
					Link:        "three",
				},
			}
			expectedQueryArgs := make([]driver.Value, 0, len(quiz)*3)

			for _, req := range quiz {
				expectedQueryArgs = append(expectedQueryArgs, req.UserId, req.ClassroomId, req.Link)
			}

			expectedNewIds := []uint64{1, 2, 3, 4}
			returnRows := sqlmock.NewRows([]string{"id"})
			for _, id := range expectedNewIds {
				returnRows.AddRow(id)
			}

			dbMock.ExpectPrepare(
				"INSERT INTO quiz \\(user_id,classroom_id,link\\) VALUES \\(\\$1,\\$2,\\$3\\),\\(\\$4,\\$5,\\$6\\),\\(\\$7,\\$8,\\$9\\) RETURNING id",
			).
				ExpectQuery().
				WithArgs(expectedQueryArgs...).
				WillReturnRows(returnRows)

			ids, err := rep.AddEntities(ctx, quiz)
			Expect(err).ToNot(HaveOccurred())
			Expect(ids).To(Equal(expectedNewIds))
		})
		It("ListEntities", func() {
			dbRows := [][]driver.Value{
				{uint64(1), uint64(10), uint64(100), "one"},
				{uint64(2), uint64(20), uint64(200), "two"},
				{uint64(3), uint64(30), uint64(300), "three"},
			}
			expectedRequests := make([]models.Quiz, 0, len(dbRows))
			returnRows := sqlmock.NewRows([]string{"id", "user_id", "classroom_id", "link"})

			for _, row := range dbRows {
				expectedRequests = append(expectedRequests, models.Quiz{
					Id:          row[0].(uint64),
					UserId:      row[1].(uint64),
					ClassroomId: row[2].(uint64),
					Link:        row[3].(string),
				})
				returnRows.AddRow(row...)
			}

			offset, limit := uint64(1), uint64(10)

			dbMock.ExpectPrepare(
				"SELECT id, user_id, classroom_id, link FROM quiz LIMIT 10 OFFSET 1",
			).
				ExpectQuery().
				WillReturnRows(returnRows)
			actualRequests, err := rep.ListEntities(ctx, limit, offset)
			Expect(err).ToNot(HaveOccurred())

			Expect(actualRequests).To(Equal(expectedRequests))
		})
		It("Remove if exists", func() {
			reqId := uint64(100)
			res := sqlmock.NewResult(0, 1)

			dbMock.ExpectPrepare(
				"DELETE FROM quiz WHERE id=\\$1",
			).
				ExpectExec().
				WithArgs(reqId).
				WillReturnResult(res)

			found := rep.RemoveEntity(ctx, reqId)
			Expect(found).To(Equal(true))
		})
		It("Remove if not exists", func() {
			reqId := uint64(100)

			dbMock.ExpectPrepare(
				"DELETE FROM quiz WHERE id=\\$1",
			).
				ExpectExec().
				WithArgs(reqId)

			found := rep.RemoveEntity(ctx, reqId)
			Expect(found).To(Equal(false))
		})
		It("Select one when if exists", func() {
			reqId := uint64(1)
			expectedReq := models.Quiz{
				Id:          1,
				ClassroomId: 1,
				UserId:      1,
				Link:        "one",
			}

			returnRows := sqlmock.
				NewRows([]string{"id", "user_id", "classroom_id", "link"}).
				AddRow(expectedReq.Id, expectedReq.UserId, expectedReq.ClassroomId, expectedReq.Link)

			dbMock.ExpectPrepare(
				"SELECT id, user_id, classroom_id, link FROM quiz WHERE id=\\$1",
			).
				ExpectQuery().
				WithArgs(reqId).
				WillReturnRows(returnRows)

			actualReq, err := rep.DescribeEntity(ctx, reqId)
			Expect(err).ToNot(HaveOccurred())
			Expect(actualReq).To(Equal(&expectedReq))
		})

		It("Select one when if not exists", func() {
			reqId := uint64(1)

			returnRows := sqlmock.
				NewRows([]string{"id", "user_id", "type", "text"})

			dbMock.ExpectPrepare(
				"SELECT id, user_id, classroom_id, link FROM quiz WHERE id=\\$1",
			).
				ExpectQuery().
				WithArgs(reqId).
				WillReturnRows(returnRows)

			actualReq, err := rep.DescribeEntity(ctx, reqId)
			Expect(err).To(Equal(NotExists))
			Expect(actualReq).To(BeNil())
		})

		It("Update exists entity", func() {
			quiz := models.Quiz{
				Id:          1,
				ClassroomId: 1,
				UserId:      1,
				Link:        "one",
			}

			res := sqlmock.NewResult(0, 1)

			dbMock.ExpectPrepare(
				"UPDATE quiz SET classroom_id = \\$1, user_id = \\$2, link = \\$3 WHERE id=\\$4",
			).
				ExpectExec().
				WithArgs(quiz.ClassroomId, quiz.UserId, quiz.Link, quiz.Id).
				WillReturnResult(res)

			updated, err := rep.UpdateEntity(ctx, quiz)
			Expect(updated).To(Equal(true))
			Expect(err).ToNot(HaveOccurred())
		})

		It("Update not-exists entity", func() {
			quiz := models.Quiz{
				Id:          1,
				ClassroomId: 1,
				UserId:      1,
				Link:        "one",
			}

			res := sqlmock.NewResult(0, 0)

			dbMock.ExpectPrepare(
				"UPDATE quiz SET classroom_id = \\$1, user_id = \\$2, link = \\$3 WHERE id=\\$4",
			).
				ExpectExec().
				WithArgs(quiz.ClassroomId, quiz.UserId, quiz.Link, quiz.Id).
				WillReturnResult(res)

			updated, err := rep.UpdateEntity(ctx, quiz)
			Expect(updated).To(Equal(false))
			Expect(err).ToNot(HaveOccurred())
		})

		Context("DB errors", func() {
			expectedError := errors.New("some error")
			It("on ListEntites", func() {

				offset, limit := uint64(10), uint64(1)

				dbMock.ExpectPrepare(
					"SELECT id, user_id, classroom_id, link FROM quiz LIMIT 1 OFFSET 10",
				).
					ExpectQuery().
					WillReturnError(expectedError)
				_, err := rep.ListEntities(ctx, limit, offset)
				Expect(err).To(Equal(expectedError))
			})

			It("on AddEntity", func() {

				newReq := models.Quiz{
					Id:          0,
					UserId:      1,
					ClassroomId: 2,
					Link:        "one",
				}
				dbMock.ExpectPrepare(
					"INSERT INTO quiz \\(user_id,classroom_id,link\\) VALUES \\(\\$1,\\$2,\\$3\\) RETURNING id",
				).
					ExpectQuery().
					WithArgs(newReq.UserId, newReq.ClassroomId, newReq.Link).
					WillReturnError(expectedError)

				newId, err := rep.AddEntity(ctx, newReq)
				Expect(err).To(Equal(expectedError))
				Expect(newId).To(Equal(uint64(0)))

			})

			It("on AddEntites", func() {

				quiz := models.Quiz{
					Id:          0,
					UserId:      1,
					ClassroomId: 2,
					Link:        "one",
				}
				expectedError := errors.New("test")
				dbMock.ExpectPrepare(
					"INSERT INTO quiz \\(user_id,classroom_id,link\\) VALUES \\(\\$1,\\$2,\\$3\\) RETURNING id",
				).
					ExpectQuery().
					WithArgs(quiz.UserId, quiz.ClassroomId, quiz.Link).
					WillReturnError(expectedError)

				ids, err := rep.AddEntities(ctx, []models.Quiz{quiz})
				Expect(err).To(Equal(expectedError))
				Expect(ids).To(BeNil())

			})
		})
	})
})
