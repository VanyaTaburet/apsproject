package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vladimirfedunov/2chan-clone/internal/entity"
)

type ThreadWithOP struct {
	Thread *entity.Thread
	OP     *entity.Post
}

func (h *Handler) AdminLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_login.html", nil)
}

func (h *Handler) AdminLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	admin, err := h.AdminService.Login(c.Request.Context(), username, password)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "admin_login.html", gin.H{"Error": err.Error()})
		return
	}
	c.SetCookie("admin_id", fmt.Sprintf("%d", admin.ID()), 3600*24, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/admin")
}

func (h *Handler) AdminLogout(c *gin.Context) {
	c.SetCookie("admin_id", "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/admin/login")
}

func (h *Handler) AdminPanel(c *gin.Context) {
	threads, err := h.ThreadService.GetAllThreads(c.Request.Context())
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin_panel.html", gin.H{"Error": err.Error()})
		return
	}
	var threadsWithOP []ThreadWithOP
	for _, t := range threads {
		posts, _ := h.PostService.ListPosts(c.Request.Context(), t.ID(), 1, 0)
		var op *entity.Post
		if len(posts) > 0 {
			op = posts[0]
		}
		threadsWithOP = append(threadsWithOP, ThreadWithOP{Thread: t, OP: op})
	}
	c.HTML(http.StatusOK, "admin_panel.html", gin.H{"Threads": threadsWithOP})
}

func (h *Handler) AdminDeleteThread(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "invalid id")
		return
	}
	if err := h.ThreadService.DeleteThread(c.Request.Context(), id); err != nil {
		c.String(500, err.Error())
		return
	}
	c.Redirect(303, "/admin")
}

func (h *Handler) AdminStickyThread(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "invalid id")
		return
	}
	thread, _, err := h.ThreadService.GetThread(c.Request.Context(), id)
	if err != nil {
		c.String(404, "thread not found")
		return
	}
	isSticky := thread.Sticky()
	if err := h.ThreadService.StickyThread(c.Request.Context(), id, !isSticky); err != nil {
		c.String(500, err.Error())
		return
	}
	c.Redirect(303, "/admin")
}

func (h *Handler) AdminLockThread(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "invalid id")
		return
	}
	thread, _, err := h.ThreadService.GetThread(c.Request.Context(), id)
	if err != nil {
		c.String(404, "thread not found")
		return
	}
	isLocked := thread.Locked()
	if err := h.ThreadService.LockThread(c.Request.Context(), id, !isLocked); err != nil {
		c.String(500, err.Error())
		return
	}
	c.Redirect(303, "/admin")
}
