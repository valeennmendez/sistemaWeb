package main

import (
	//"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/valeennmendez/api-go/connection"
	"github.com/valeennmendez/api-go/models"
	"github.com/valeennmendez/api-go/routes"
)

func main() {
	connection.ConnectionDB()

	connection.DB.AutoMigrate(&models.Patients{})
	connection.DB.AutoMigrate(&models.Admin{})
	connection.DB.AutoMigrate(&models.Appoinment{})

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://sistemaweb-production.up.railway.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

 	// Servir archivos estáticos
	r.Static("/static", "./static")
	r.StaticFile("/login.html", "./pages/login.html") 
	r.StaticFile("/index.html", "./pages/index.html") 

	// Rutas de autenticación
	r.POST("/register", routes.RegisterUser)
	r.POST("/login", routes.Login)
	r.GET("/validate", routes.ValidateSession)
	r.POST("/logout", routes.CloseSesion)

	// Rutas protegidas
	protected := r.Group("/")
	protected.Use(routes.AuthMiddleware())
	{
		protected.StaticFile("/index.html", "./pages/index.html")
		protected.GET("/patients", routes.GetAllPatients)
		protected.GET("/patients/:id", routes.GetPatientByID)
		protected.POST("/create", routes.CreatePatient)
		protected.PUT("/edit/:id", routes.EditPatient)
		protected.DELETE("/patients/:id", routes.DeletePacients)
		protected.GET("/total-patients", routes.TotalPatientsData)

	}

	r.POST("/create-appointment", routes.CreateAppoinment)
	r.GET("/appointment-today", routes.AppointmentToday)
	r.GET("/available-hours", routes.GetAviableHours)
	r.GET("/search-patient", routes.SearchPatient)
	r.GET("/appointments", routes.GetAllAppointments)

	
	r.DELETE("/cancel-appointment/:id", routes.CancelAppointment) // <--- DEBE ESTAR PUBLICA SI O SI.

	// Ruta raíz
/* 	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusOK, "/login.html")
	})
 */
	r.Run("0.0.0.0:8080")
}
