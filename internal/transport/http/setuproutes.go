package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vladimirfedunov/2chan-clone/internal/usecase"
)

func SetupRoutes(router *gin.Engine, h *Handler, adminService *usecase.AdminService) {
	router.Static("/static", "./web")

	// Main pages
	router.GET("/", h.IndexPage)
	router.GET("/boards/:slug", h.BoardPage)
	router.GET("/about", h.AboutPage) // Новый роут

	// Boards
	router.GET("/boards", h.GetBoards)

	// Threads
	router.GET("/boards/:slug/threads", h.GetThreads)
	router.POST("/boards/:slug/threads", h.CreateThread)
	router.GET("/threads/:id", h.GetThread)
	router.PATCH("/threads/:id/sticky", h.StickyThread)
	router.PATCH("/threads/:id/lock", h.LockThread)
	router.DELETE("/threads/:id", h.DeleteThread)

	// Posts
	router.GET("/threads/:id/posts", h.GetPosts)
	router.POST("/threads/:id/posts", h.CreatePost)
	router.GET("/posts/:id", h.GetPost)
	router.DELETE("/posts/:id", h.DeletePost)

	// Admin
	router.GET("/admin/login", h.AdminLoginPage)
	router.POST("/admin/login", h.AdminLogin)
	router.GET("/admin/logout", h.AdminLogout)
	router.GET("/admin", adminAuthMiddleware(adminService), h.AdminPanel)
	router.POST("/admin/threads/:id/delete", h.AdminDeleteThread)
	router.POST("/admin/threads/:id/sticky", h.AdminStickyThread)
	router.POST("/admin/threads/:id/lock", h.AdminLockThread)
}

func adminAuthMiddleware(adminService *usecase.AdminService) gin.HandlerFunc {
	return func(c *gin.Context) {
		adminID, err := c.Cookie("admin_id")
		if err != nil || adminID == "" {
			c.Redirect(http.StatusSeeOther, "/admin/login")
			c.Abort()
			return
		}

		var id int64
		_, err = fmt.Sscanf(adminID, "%d", &id)
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/admin/login")
			c.Abort()
			return
		}

		admin, err := adminService.GetByID(c.Request.Context(), id)
		if err != nil || admin == nil {
			c.Redirect(http.StatusSeeOther, "/admin/login")
			c.Abort()
			return
		}

		c.Next()
	}
}
