package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["pwd"]), 14)
	user := User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	var existingUser User
	DB.Where("email=?", data["email"]).First(&existingUser)

	if existingUser.Email == data["email"] {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User with provided emailId already registerd. Please login to proceed.",
		})
	} else {
		DB.Create(&user)

		return c.JSON(user)
	}

}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user User

	DB.Where("email=?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["pwd"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func UserSession(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user User

	DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Participate(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	// cookie := c.Cookies("jwt")

	// token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
	// 	return []byte(SecretKey), nil
	// })

	// if err != nil {
	// 	c.Status(fiber.StatusUnauthorized)
	// 	return c.JSON(fiber.Map{
	// 		"message": "unauthenticated",
	// 	})
	// }

	// claims := token.Claims.(*jwt.StandardClaims)

	// var exitingLottery Lottery

	// DB.Where("id = ?", claims.Issuer).First(&exitingLottery)

	var user User
	DB.Where("email = ?", data["participantEmail"]).First(&user)

	if user.Email != data["participantEmail"] {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": "Register the user first",
		})
	}

	var existingParticipants []Participant

	DB.Where("lottery_id = ?", data["lotteryId"]).Or("participant_email = ?", data["participantEmail"]).Find(&existingParticipants)

	temp, err := strconv.ParseUint(data["lotteryId"], 10, 64)

	if err != nil {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": "Type conversion failed",
		})
	}

	lotteryId := uint(temp)

	for _, participant := range existingParticipants {
		if participant.LotteryId == lotteryId && participant.ParticipantEmail == data["participantEmail"] {
			c.Status(fiber.StatusConflict)
			return c.JSON(fiber.Map{
				"message": "Participant has already opted for this lottery.",
			})
		}
	}

	participant := Participant{
		LotteryId:        lotteryId,
		ParticipantEmail: data["participantEmail"],
	}

	var lottery Lottery

	DB.Where("id=?", lotteryId).First(&lottery)

	if lottery.Participants < lottery.Limit {
		DB.Create(&participant)
		newCount := lottery.Participants + 1

		DB.Model(&lottery).Update("participants", newCount)
	} else if lottery.Participants == lottery.Limit {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": "Lottery is closed for participation. Winner will be disclosed soon. You will be notified on opening of new lottery.",
		})
	}

	return c.JSON(participant)
}

func CreateLottery(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	tempLim, err := strconv.ParseInt(data["limit"], 10, 64)
	if err != nil {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": "Type conversion failed",
		})
	}

	limit := int(tempLim)

	lottery := Lottery{
		LotteryName: data["lotteryName"],
		Limit:       limit,
	}

	var existingLottery Lottery

	DB.Where("lottery_name=?", data["lotteryName"]).First(&existingLottery)

	if existingLottery.LotteryName == lottery.LotteryName {
		c.Status(fiber.StatusBadGateway)
		return c.JSON(fiber.Map{
			"message": "Lottery already created",
		})
	} else {
		DB.Create(&lottery)

		return c.JSON(lottery)
	}

}

func ChooseWinner(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var lottery Lottery
	DB.Where("id=?", data["lotteryId"]).First(&lottery)

	var winner Winner
	if lottery.Winner == 0 {
		if lottery.Limit == lottery.Participants {

			myRandomNum, err := rand.Int(rand.Reader, big.NewInt(int64(lottery.Limit)))

			for int(myRandomNum.Int64()) <= 0 {
				myRandomNum, err = rand.Int(rand.Reader, big.NewInt(int64(lottery.Limit)))

			}
			if err != nil {
				c.Status(fiber.StatusConflict)
				return c.JSON(fiber.Map{
					"message": "Type conversion failed",
				})
			}
			var user User

			DB.Where("id = ?", int(myRandomNum.Int64())).First(&user)

			// TODO: Generate radom number based on no of participants
			fmt.Printf("Generated Random no %v and Type %T\n", myRandomNum, myRandomNum)
			DB.Model(&lottery).Update("winner", user.Id)
			winner = Winner{
				LotteryId:   lottery.Id,
				WinnerEmail: user.Email,
			}
			DB.Create(&winner)

		} else if lottery.Limit > lottery.Participants {
			c.Status(fiber.StatusBadGateway)
			return c.JSON(fiber.Map{

				"message": "Lottery is still open and waiting for more participants",
			})
		}
	} else if lottery.Winner != 0 {

		var user User
		DB.Where("id = ?", lottery.Winner).First(&user)
		winMsg := "Winner already opted Lottery winner is " + user.Email

		c.Status(fiber.StatusBadGateway)
		return c.JSON(fiber.Map{

			"message": winMsg,
		})

	}

	return c.JSON(&winner)

}
