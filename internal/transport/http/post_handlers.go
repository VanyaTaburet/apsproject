package http

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPosts(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid thread id"})
		return
	}
	posts, err := h.PostService.ListPosts(c.Request.Context(), id, 50, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h *Handler) CreatePost(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid thread id")
		return
	}
	var req struct {
		BoardSlug string  `form:"board_slug"`
		Author    string  `form:"author"`
		Text      string  `form:"text"`
		ImageURL  *string `form:"image_url"`
		ParentID  *int64  `form:"parent_id"`
		Tripcode  *string `form:"tripcode"`
	}
	if err := c.ShouldBind(&req); err != nil {
		c.String(http.StatusBadRequest, "invalid request")
		return
	}
	// Обработка картинки
	file, err := c.FormFile("image")
	var imageURL *string
	if err == nil {
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
		path := "web/uploads/" + filename
		if err := c.SaveUploadedFile(file, path); err == nil {
			url := "/static/uploads/" + filename
			imageURL = &url
		}
	}
	ip := c.ClientIP()
	_, err = h.PostService.CreatePost(c.Request.Context(), id, req.BoardSlug, req.Author, req.Text, imageURL, req.ParentID, req.Tripcode, ip)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Redirect(http.StatusSeeOther, "/threads/"+strconv.FormatInt(id, 10))
}

func (h *Handler) GetPost(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}
	post, err := h.PostService.GetPost(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *Handler) DeletePost(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}
	if err := h.PostService.DeletePost(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
