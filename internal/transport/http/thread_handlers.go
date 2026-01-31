package http

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetThreads(c *gin.Context) {
	slug := c.Param("slug")
	threads, err := h.ThreadService.ListThreads(c.Request.Context(), slug, 20, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, threads)
}

func (h *Handler) CreateThread(c *gin.Context) {
	slug := c.Param("slug")
	var req struct {
		Author   string  `form:"author"`
		Title    string  `form:"title"`
		Text     string  `form:"text"`
		ImageURL *string `form:"image_url"`
		Tripcode *string `form:"tripcode"`
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
	thread, err := h.ThreadService.CreateThread(c.Request.Context(), slug, req.Author, req.Title, req.Text, imageURL, req.Tripcode, ip)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if c.GetHeader("HX-Request") != "" {
		c.Header("HX-Redirect", "/threads/"+strconv.FormatInt(thread.ID(), 10))
		c.Status(http.StatusNoContent)
		return
	}
	c.Redirect(http.StatusSeeOther, "/threads/"+strconv.FormatInt(thread.ID(), 10))
}

func (h *Handler) GetThread(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid id")
		return
	}
	thread, posts, err := h.ThreadService.GetThread(c.Request.Context(), id)
	if err != nil {
		c.String(http.StatusNotFound, "thread not found")
		return
	}
	c.HTML(http.StatusOK, "thread.html", gin.H{
		"Thread": thread,
		"Posts":  posts,
	})
}

func (h *Handler) StickyThread(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req struct {
		Sticky bool `json:"sticky"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if err := h.ThreadService.StickyThread(c.Request.Context(), id, req.Sticky); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) LockThread(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req struct {
		Locked bool `json:"locked"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if err := h.ThreadService.LockThread(c.Request.Context(), id, req.Locked); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) DeleteThread(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.ThreadService.DeleteThread(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
