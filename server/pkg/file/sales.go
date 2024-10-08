package file

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/server/pkg/notification"
	"github.com/shahsuman438/SALES-DASH/server/pkg/sales"
	"github.com/shahsuman438/SALES-DASH/server/pkg/utils/logger"
)

func ProcessSalesFile(pathToCsv string) error {
	file, err := os.Open(pathToCsv)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	if _, err := reader.Read(); err != nil {
		logger.Error("Error reading header", err)
		return err
	}

	var transactions []sales.Sales
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error("Error reading record:", err)
			return err
		}
		var transaction sales.Sales
		tId, err := strconv.Atoi(record[0])
		if err != nil {
			logger.Error("Error Converting tId", err)
			continue
		}
		transaction.TransactionId = tId
		pId, err := strconv.Atoi(record[1])
		if err != nil {
			logger.Error("Error Converting TransactionId", err)
			continue
		}
		transaction.ProductId = pId
		qty, err := strconv.Atoi(record[2])
		if err != nil {
			logger.Error("Error Converting productId", err)
			continue
		}
		transaction.Quantity = qty

		tAmount, err := strconv.Atoi(record[3])
		if err != nil {
			logger.Error("Error Converting transaction amount", err)
			continue
		}

		transaction.TotalTransactionAmount = float64(tAmount)
		transaction.TransactionDate = record[4]

		transactions = append(transactions, transaction)
	}

	err = sales.AddMany(&gin.Context{}, &transactions)
	if err != nil {
		logger.Error("Unable to add data in Sales", err)
	}
	notification.SendMessageToClients("New File Processed, Please re-sync the Table.")
	logger.Info(fmt.Sprintf("PROCESSED File %s ", pathToCsv))
	return nil
}
