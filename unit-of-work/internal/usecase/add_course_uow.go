package usecase

import (
	"context"

	"github.com/felipemagrassi/go-expert/unit-of-work/internal/entity"
	"github.com/felipemagrassi/go-expert/unit-of-work/internal/repository"
	"github.com/felipemagrassi/go-expert/unit-of-work/pkg/uow"
)

type InputUseCaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCaseUow struct {
	Uow uow.UowInterface
}

func NewAddCourseUseCaseUow(uow uow.UowInterface) *AddCourseUseCaseUow {
	return &AddCourseUseCaseUow{
		Uow: uow,
	}
}

func (a *AddCourseUseCaseUow) Execute(ctx context.Context, input InputUseCase) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		category := entity.Category{
			Name: input.CategoryName,
		}

		categoryRepo, err := a.GetCategoryRepository(ctx)
		if err != nil {
			return err
		}

		err = categoryRepo.Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}

		courseRepo, err := a.GetCourseRepository(ctx)
		if err != nil {
			return err
		}

		err = courseRepo.Insert(ctx, course)
		if err != nil {
			return err
		}

		return nil
	})
}

func (a *AddCourseUseCaseUow) GetCategoryRepository(ctx context.Context) (repository.CategoryRepositoryInterface, error) {
	r, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		return nil, err
	}
	if categoryRepository, ok := r.(repository.CategoryRepositoryInterface); ok {
		return categoryRepository, nil
	}
	return nil, err
}

func (a *AddCourseUseCaseUow) GetCourseRepository(ctx context.Context) (repository.CourseRepositoryInterface, error) {
	r, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		return nil, err
	}
	if courseRepository, ok := r.(repository.CourseRepositoryInterface); ok {
		return courseRepository, nil
	}
	return nil, err
}
