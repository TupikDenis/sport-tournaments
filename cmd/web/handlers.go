package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sport-tournaments/pkg/algoritms/olympic"
	"sport-tournaments/pkg/algoritms/robin_round"
	"sport-tournaments/pkg/handlers/match"
	"sport-tournaments/pkg/handlers/team"
	tournament2 "sport-tournaments/pkg/handlers/tournament"
	"sport-tournaments/pkg/handlers/user"
	"sport-tournaments/pkg/models"
	"sport-tournaments/pkg/models/databaseModels"
	"sport-tournaments/pkg/services/feedback"
	"sport-tournaments/pkg/services/token"
	"strconv"
)

var tokenStr string

func handle() {
	router := createRouter()

	homePage(router)
	feedbackPage(router)
	profilePage(router)

	userRouterGroup(router)
	tournamentRouterGroup(router)

	startServer(router, ":8080")
}

func createRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./ui/static")
	router.LoadHTMLGlob("ui/html/*")

	return router
}

func homePage(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		user, err := token.ParseToken(tokenStr)

		if err != nil {
			tokenStr = ""
		}

		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Tournament maker",
			"token": tokenStr,
			"id":    user.Id,
		})
	})
}

func feedbackPage(router *gin.Engine) {
	router.GET("/feedback", func(c *gin.Context) {
		user, err := token.ParseToken(tokenStr)

		if err != nil {
			tokenStr = ""
		}

		c.HTML(http.StatusOK, "feedback.html", gin.H{
			"title": "Feedback",
			"token": tokenStr,
			"id":    user.Id,
		})
	})

	router.POST("/feedback/send", func(c *gin.Context) {
		email := c.PostForm("email")
		text := c.PostForm("text")

		feedback.Feedback(email, text)

		user, err := token.ParseToken(tokenStr)

		if err != nil {
			tokenStr = ""
		}

		c.HTML(http.StatusOK, "feedback.html", gin.H{
			"title": "Feedback",
			"token": tokenStr,
			"id":    user.Id,
		})
	})
}

func profilePage(router *gin.Engine) {
	router.GET("/profile", func(c *gin.Context) {
		user, err := token.ParseToken(tokenStr)

		if err != nil {
			tokenStr = ""
			c.HTML(http.StatusUnauthorized, "error.html", gin.H{
				"title": "Invalid token! Please, authorize again!",
			})
			return
		}

		c.HTML(http.StatusOK, "profile.html", gin.H{
			"title":    user.Username,
			"token":    tokenStr,
			"id":       user.Id,
			"username": user.Username,
			"role":     user.Role,
		})
	})
}

func userRouterGroup(router *gin.Engine) {
	router.GET("/sign", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
		})
	})

	router.GET("/logout", func(c *gin.Context) {
		tokenStr = ""

		c.Redirect(http.StatusFound, "/")
	})

	router.POST("/authentication", func(c *gin.Context) {
		username := c.PostForm("username-sign-in")
		password := c.PostForm("password-sign-in")

		tokenStr = user.Login(username, password)

		c.Redirect(http.StatusFound, "/")
	})

	users := router.Group("/api/users")
	{
		users.POST("", func(c *gin.Context) {
			username := c.PostForm("username-sign-up")
			password := c.PostForm("password-sign-up")

			user.CreateUser(username, password)
			c.Redirect(http.StatusFound, "/sign")
		})

		users.GET("/:id", func(c *gin.Context) {

			c.Redirect(http.StatusFound, "/profile")
		})

		users.PATCH("/password/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.ParseUint(idStr, 10, 32)

			if err != nil {
				panic(err)
			}

			oldPasswordForm := c.PostForm("old_password")
			oldPassword := user.GetUserPassword(uint(id))

			if oldPassword != oldPasswordForm {
				c.Redirect(http.StatusNotAcceptable, "/profile")
				return
			}

			newPassword := c.PostForm("new_password")

			user.UpdateUserPassword(uint(id), newPassword)
		})
	}
}

func tournamentRouterGroup(router *gin.Engine) {
	router.GET("/tournament/create", func(c *gin.Context) {
		user, err := token.ParseToken(tokenStr)

		if err != nil {
			tokenStr = ""
			c.HTML(http.StatusUnauthorized, "error.html", gin.H{
				"title": "Invalid token! Please, authorize again!",
			})
			return
		}

		c.HTML(http.StatusOK, "tournament_form.html", gin.H{
			"title": "Create tournament",
			"token": tokenStr,
			"id":    user.Id,
		})
	})

	router.GET("/tournament", func(c *gin.Context) {
		user, _ := token.ParseToken(tokenStr)

		tournaments := tournament2.GetAllTournaments()

		c.HTML(http.StatusOK, "tournament.html", gin.H{
			"title":       "Tournaments",
			"token":       tokenStr,
			"tournaments": tournaments,
			"id":          user.Id,
		})
	})

	router.GET("/tournament/user/:id", func(c *gin.Context) {
		user, err := token.ParseToken(tokenStr)

		if err != nil {
			tokenStr = ""
			c.HTML(http.StatusUnauthorized, "error.html", gin.H{
				"title": "Invalid token! Please, authorize again!",
			})
			return
		}

		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)

		if err != nil {
			panic(err)
		}

		tournaments := tournament2.GetAllTournamentsByUserId(uint(id))

		c.HTML(http.StatusOK, "tournament.html", gin.H{
			"title":       "Tournaments",
			"token":       tokenStr,
			"tournaments": tournaments,
			"id":          user.Id,
		})
	})

	router.GET("/tournament/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)

		if err != nil {
			panic(err)
		}

		tournament := tournament2.GetTournamentById(uint(id))
		tournamentMatches := match.GetMatchesByIdTournament(tournament.Id)

		var pairs []models.Pair
		var matches []models.Match

		tournamentMatches = append(tournamentMatches, models.TransformedFullMatch{})

		for k := 0; k < len(tournamentMatches)-1; k++ {
			pairs = append(pairs, models.Pair{
				Number:    tournamentMatches[k].Number,
				HomeTeam:  tournamentMatches[k].HomeTeam,
				HomeScore: tournamentMatches[k].HomeScore,
				AwayTeam:  tournamentMatches[k].AwayTeam,
				AwayScore: tournamentMatches[k].AwayScore,
			})

			if tournamentMatches[k+1].Round != tournamentMatches[k].Round {
				matches = append(matches, models.Match{
					TournamentId: tournament.Id,
					Round:        tournamentMatches[k].Round,
					Pair:         pairs,
				})

				pairs = nil
			}
		}

		user, _ := token.ParseToken(tokenStr)

		matchesJSON, _ := json.Marshal(matches)
		fmt.Println(string(matchesJSON))

		c.HTML(http.StatusOK, "tournament_id.html", gin.H{
			"title":                   "Tournaments",
			"token":                   tokenStr,
			"tournament":              tournament,
			"tournament_matches":      matches,
			"id":                      user.Id,
			"tournament_matches_json": string(matchesJSON),
		})
	})

	tournament := router.Group("/api/tournaments")
	{
		tournament.POST("", func(c *gin.Context) {
			idStr := c.PostForm("id")
			id, err := strconv.ParseInt(idStr, 10, 32)

			if err != nil {
				panic(err)
			}

			name := c.PostForm("tournament_name")
			description := c.PostForm("description")
			sport := c.PostForm("sport")
			system := c.PostForm("system")

			tournament2.CreateTournament(name, description, sport, uint(id), system)
			tournamentId := tournament2.GetLastTournamentId()

			if system == "robin" {
				numberStr := c.PostForm("team_number_robin")
				number, err2 := strconv.ParseInt(numberStr, 10, 32)

				if err2 != nil {
					panic(err2)
				}

				robinStr := c.PostForm("number_robin")
				robin, err3 := strconv.ParseInt(robinStr, 10, 32)

				if err3 != nil {
					panic(err3)
				}

				mixedTeamRobin := c.PostForm("mixed_team_robin")
				var mixed bool = false

				if mixedTeamRobin == "mixed_team_robin" {
					mixed = true
				}

				var teamName []string
				var teams []databaseModels.TransformedTeam

				for i := 0; i < int(number); i++ {
					tmp := i + 1
					teamName = append(teamName, c.PostForm("team_name"+strconv.Itoa(tmp)+"_robin"))
					team.CreateTeam(teamName[i], tournamentId)
					teams = append(teams, databaseModels.TransformedTeam{
						TeamName: teamName[i],
					})
				}

				matches := robin_round.RobinRound(teams, int(robin), mixed, tournamentId)
				match.CreateSchedule(matches, int(robin))
			}

			if system == "knockout" {
				numberStr := c.PostForm("team_number_knockout")
				number, err2 := strconv.ParseInt(numberStr, 10, 32)

				if err2 != nil {
					panic(err2)
				}

				var teamName []string
				var teams []databaseModels.TransformedTeam

				for i := 0; i < int(number); i++ {
					tmp := i + 1
					teamName = append(teamName, c.PostForm("team_name"+strconv.Itoa(tmp)+"_knockout"))
					team.CreateTeam(teamName[i], tournamentId)
					teams = append(teams, databaseModels.TransformedTeam{
						TeamName: teamName[i],
					})
				}

				mixedTeamKnockout := c.PostForm("mixed_team_knockout")
				var mixed bool = false

				if mixedTeamKnockout == "mixed_team_knockout" {
					mixed = true
				}

				matches := olympic.Olympic(teams, mixed, tournamentId)
				match.CreateSchedule(matches, 0)
			}

			c.Redirect(http.StatusFound, "/tournament")
		})

		tournament.GET("", func(c *gin.Context) {
			c.Redirect(http.StatusFound, "/tournament")
		})

		tournament.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.Redirect(http.StatusFound, "/tournament/"+id)
		})

		tournament.DELETE("/:id", func(c *gin.Context) {
			idStr := c.Param("id")
			id, err := strconv.ParseUint(idStr, 10, 32)

			if err != nil {
				panic(err)
			}

			tournament2.DeleteTournament(uint(id))
		})
	}
}

func startServer(router *gin.Engine, port string) {
	err := router.Run(port)

	if err != nil {
		panic(err)
	}
}
