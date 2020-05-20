package forms

type AddTagForm struct {
    Name string `form:"name" valid:"Required;MaxSize(20)"`
    CreatedBy string `form:"create_by"`
    State int `form:"state"`
}

type TagListForm1 struct {
    Name string `form:"name" json:"name" binding:"required"`
    Email string `form:"email" json:"email" binding:"required,email"`
    Page int `form:"page" json:"page"`
}
type TagListForm2 struct {
    Name string `form:"name" validate:"required"`
    Email string `form:"email" validate:"required,email"`
    Page int `form:"page" json:"page"`
}

type TagDeatilForm struct {
    Id int `form:"id" json:"id" binding:"required"`
}