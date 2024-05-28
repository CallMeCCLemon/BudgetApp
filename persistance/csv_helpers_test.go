package persistance

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestParseCsvTransactions(t *testing.T) {
	AccountID := uint(47)

	type args struct {
		filepath string
	}
	tests := []struct {
		name             string
		args             args
		wantTransactions []Transaction
		wantErr          assert.ErrorAssertionFunc
	}{
		{
			name: "Single Record Test",
			args: args{
				filepath: "fixtures/singleTransaction.csv",
			},
			wantTransactions: []Transaction{
				{
					Date:       time.Date(2024, 05, 22, 0, 0, 0, 0, time.UTC),
					Amount:     -16.01,
					Memo:       "Test Transaction",
					AccountID:  AccountID,
					CategoryID: 2,
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Multi-Record Test",
			args: args{
				filepath: "fixtures/multipleTransactions.csv",
			},
			wantTransactions: []Transaction{
				{
					Date:       time.Date(2024, 05, 22, 0, 0, 0, 0, time.UTC),
					Amount:     -16.01,
					Memo:       "payment-to-something",
					AccountID:  AccountID,
					CategoryID: 2,
				},
				{
					Date:       time.Date(2024, 05, 20, 0, 0, 0, 0, time.UTC),
					Amount:     -70.88,
					Memo:       "venmo-payment",
					AccountID:  AccountID,
					CategoryID: 4,
				},
				{
					Date:       time.Date(2024, 05, 20, 0, 0, 0, 0, time.UTC),
					Amount:     -365.27,
					Memo:       "amex-payment",
					AccountID:  AccountID,
					CategoryID: 4,
				},
				{
					Date:       time.Date(2024, 05, 15, 0, 0, 0, 0, time.UTC),
					Amount:     12345.67,
					Memo:       "PAYROLL",
					AccountID:  AccountID,
					CategoryID: 4,
				},
				{
					Date:       time.Date(2024, 05, 14, 0, 0, 0, 0, time.UTC),
					Amount:     -4.13,
					Memo:       "paypal",
					AccountID:  AccountID,
					CategoryID: 2,
				},
				{
					Date:       time.Date(2024, 05, 14, 0, 0, 0, 0, time.UTC),
					Amount:     -365.31,
					Memo:       "amazon",
					AccountID:  AccountID,
					CategoryID: 6,
				},
				{
					Date:       time.Date(2024, 05, 13, 0, 0, 0, 0, time.UTC),
					Amount:     43.85,
					Memo:       "venmo-cashout",
					AccountID:  AccountID,
					CategoryID: 6,
				},
				{
					Date:       time.Date(2024, 05, 8, 0, 0, 0, 0, time.UTC),
					Amount:     1234.00,
					Memo:       "tax-refund",
					AccountID:  AccountID,
					CategoryID: 2,
				},
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Open(tt.args.filepath)
			if err != nil {
				t.Error(err)
			}
			gotTransactions, err := ParseCsvTransactions(file, AccountID)
			if !tt.wantErr(t, err, fmt.Sprintf("ParseCsvTransactions(%v)", tt.args.filepath)) {
				return
			}
			assert.Equalf(t, tt.wantTransactions, gotTransactions, "ParseCsvTransactions(%v)", tt.args.filepath)
		})
	}
}
