package functions

//imports
import (
	"Conveter/tables"
	"net/http"

	"github.com/gin-gonic/gin"
)

// slice of tables with data from db
type Currencies struct {
	Currcs []tables.Currency
}

// declare all functions
type Functions interface {
	ShowAll()
	ShowOne(code string) (*tables.Currency, error)
	Change(id1, id2 int) (float32, error)
}

// write function ShowAll
func (c *Currencies) ShowAll(ctx *gin.Context) {
	ctx.JSON(200, c.Currcs)

	if c.Currcs == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"errorcode": 400,
		})
	}

}

// function ShowOne
func (c Currencies) ShowOne(code string, ctx *gin.Context) {
	existing := false
	for i := range len(c.Currcs) {
		if code == c.Currcs[i].Code {
			ctx.JSON(http.StatusOK, gin.H{
				"code":     c.Currcs[i].Code,
				"fullname": c.Currcs[i].FullName,
				"sign":     c.Currcs[i].Sign,
				"todollar": c.Currcs[i].ToDollar,
			})
			existing = true
		}

	}
	if !existing {
		ctx.JSON(http.StatusOK, gin.H{
			"errorCode": 401,
		})
	}
}

func (c Currencies) Change(id1, id2 int, ctx *gin.Context) {
	var rate1 float32 = 0
	var rate2 float32 = 0
	for i := range len(c.Currcs) {
		if i == id1 {
			rate1 = c.Currcs[i].ToDollar
		}
		if i == id2 {
			rate2 = c.Currcs[i].ToDollar
		}
	}
	if rate1 == 0 || rate2 == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"errorCode": 300,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"changed": rate1 / rate2,
		})
	}

}

// declaring object
var CurrObj Currencies
