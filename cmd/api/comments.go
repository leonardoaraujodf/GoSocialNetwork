package main

import (
	"net/http"

	"github.com/leonardoaraujodf/social/internal/store"
)

type createCommentPayload struct {
	Content string `json:"content" validate:"required,max=500"`
}

// CreateComment godoc
//
//	@Summary		Creates a new comment
//	@Description	Creates a new comment on a post
//	@Tags			comments
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		createCommentPayload	true	"Comment payload"
//	@Success		201		{object}	store.Comment
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/posts/{postID}/comments [post]
func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	var payload createCommentPayload
	if err := ReadJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	comment := &store.Comment{
		PostID:  post.ID,
		Content: payload.Content,
		//TODO: Get user ID from auth token
		UserID: 1,
	}

	ctx := r.Context()
	if err := app.store.Comment.Create(ctx, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
