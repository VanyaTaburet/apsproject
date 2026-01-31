package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBoards(c *gin.Context) {
	boards, err := h.BoardService.ListBoards(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, boards)
}

func (h *Handler) GetBoard(c *gin.Context) {
	slug := c.Param("slug")
	board, err := h.BoardService.GetBoard(c.Request.Context(), slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "board not found"})
		return
	}
	c.JSON(http.StatusOK, board)
}

func (h *Handler) BoardPage(c *gin.Context) {
	boardSlug := c.Param("slug")

	previews, err := h.ThreadService.ListThreadPreviews(c.Request.Context(), boardSlug)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "board.html", gin.H{"Error": err.Error()})
		return
	}

	board, err := h.BoardService.GetBoard(c.Request.Context(), boardSlug)
	if err != nil {
		c.String(http.StatusNotFound, "Доска не найдена")
		return
	}

	c.HTML(http.StatusOK, "board.html", gin.H{
		"Board":          board,
		"ThreadPreviews": previews,
	})
}
