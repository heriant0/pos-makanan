package utility

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GenerateInvoiceId(sequence *string) (res string) {
	currentTime := time.Now()
	year := currentTime.Format("20060102")

	num := 1
	if sequence != nil || *sequence != "" {
		seq := strings.Split(*sequence, "/")

		num, _ = strconv.Atoi(seq[len(seq)-1])
		num++
	}

	res = fmt.Sprintf("POSMAK/INV/%s/%d", year, num)
	return
}
