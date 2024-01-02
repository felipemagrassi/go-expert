package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"

	"github.com/felipemagrassi/go-expert/graphql/graph/model"
)

// Courses is the resolver for the courses field.
func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	courses, err := r.CourseDB.FindByCategoryID(obj.ID)
	if err != nil {
		return nil, err
	}
	coursesList := []*model.Course{}
	for _, course := range courses {
		coursesList = append(coursesList, &model.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
		})
	}

	return coursesList, nil
}

// Category is the resolver for the category field.
func (r *courseResolver) Category(ctx context.Context, obj *model.Course) (*model.Category, error) {
	category, err := r.CategoryDB.FindByCourseID(obj.ID)
	if err != nil {
		return nil, err
	}

	return &model.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: &category.Description,
	}, nil
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category, err := r.CategoryDB.Create(input.Name, *input.Description)
	if err != nil {
		return nil, err
	}

	return &model.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: &category.Description,
	}, nil
}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	course, err := r.CourseDB.Create(input.Name, *input.Description, input.CategoryID)
	if err != nil {
		return nil, err
	}

	return &model.Course{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
	}, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	category, err := r.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	categoryList := []*model.Category{}

	for _, category := range category {
		categoryList = append(categoryList, &model.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: &category.Description,
		})
	}

	return categoryList, nil
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	courses, err := r.CourseDB.FindAll()
	if err != nil {
		return nil, err
	}

	var courseList []*model.Course

	for _, course := range courses {
		courseList = append(courseList, &model.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
		})
	}

	return courseList, nil
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Course returns CourseResolver implementation.
func (r *Resolver) Course() CourseResolver { return &courseResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type courseResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
