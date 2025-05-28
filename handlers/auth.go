package handlers

import (
    "database/sql"
    "go-auth-api/models"
    "net/http"
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/labstack/echo/v4"
    "golang.org/x/crypto/bcrypt"
)

type Credentials struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

func Register(c echo.Context) error {
    creds := new(Credentials)
    if err := c.Bind(creds); err != nil {
        return err
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    _, err = models.DB.Exec("INSERT INTO users(name, email, password_hash) VALUES (?, ?, ?)",
        creds.Name, creds.Email, string(hashedPassword))
    if err != nil {
        return err
    }

    return c.JSON(http.StatusCreated, echo.Map{"message": "user registered"})
}

func Login(c echo.Context) error {
    creds := new(Credentials)
    if err := c.Bind(creds); err != nil {
        return err
    }

    row := models.DB.QueryRow("SELECT id, name, password_hash FROM users WHERE email = ?", creds.Email)

    var user models.User
    err := row.Scan(&user.ID, &user.Name, &user.PasswordHash)
    if err == sql.ErrNoRows || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password)) != nil {
        return echo.ErrUnauthorized
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":    user.ID,
        "name":  user.Name,
        "email": creds.Email,
        "exp":   time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, echo.Map{"token": tokenStr})
}

func Me(c echo.Context) error {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    return c.JSON(http.StatusOK, claims)
}
