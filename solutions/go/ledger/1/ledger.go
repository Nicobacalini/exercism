package ledger

import (
	"errors"
	"strconv"
	"strings"
    "fmt"
    "sort"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

// FormatLedger procesa, ordena y maqueta el libro contable de manera secuencial y limpia.
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if locale != "en-US" && locale != "nl-NL" {
		return "", errors.New("locale no soportado")
	}
	if currency != "USD" && currency != "EUR" {
		return "", errors.New("moneda no soportada")
	}

	// 1. Clonar las entradas para evitar efectos secundarios en el slice original
	entriesCopy := append([]Entry(nil), entries...)

	// 2. Ordenar usando la libreria nativa (Fecha -> Descripcion -> Valor)
	sort.SliceStable(entriesCopy, func(i, j int) bool {
		if entriesCopy[i].Date != entriesCopy[j].Date {
			return entriesCopy[i].Date < entriesCopy[j].Date
		}
		if entriesCopy[i].Description != entriesCopy[j].Description {
			return entriesCopy[i].Description < entriesCopy[j].Description
		}
		return entriesCopy[i].Change < entriesCopy[j].Change
	})

	// 3. Renderizar la cabecera correspondiente al locale
	var sb strings.Builder
	if locale == "nl-NL" {
		sb.WriteString("Datum      | Omschrijving              | Verandering  \n")
	} else {
		sb.WriteString("Date       | Description               | Change       \n")
	}

	// 4. Procesar secuencialmente cada entrada
	for _, entry := range entriesCopy {
		// Validar y formatear la fecha (esperado "YYYY-MM-DD")
		if len(entry.Date) != 10 || entry.Date[4] != '-' || entry.Date[7] != '-' {
			return "", errors.New("formato de fecha invalido")
		}
		year, month, day := entry.Date[0:4], entry.Date[5:7], entry.Date[8:10]
		
		var dateStr string
		if locale == "nl-NL" {
			dateStr = fmt.Sprintf("%s-%s-%s", day, month, year)
		} else {
			dateStr = fmt.Sprintf("%s/%s/%s", month, day, year)
		}

		// Formatear la descripcion (truncar a 25 caracteres maximo con "...")
		descStr := entry.Description
		if len(descStr) > 25 {
			descStr = descStr[:22] + "..."
		} else {
			descStr = descStr + strings.Repeat(" ", 25-len(descStr))
		}

		// Formatear el monto monetario financiero
		moneyStr := formatMoney(entry.Change, currency, locale)

		// Unificar la linea manteniendo el espaciado exacto de las columnas
		row := fmt.Sprintf("%-10s | %s | %13s\n", dateStr, descStr, moneyStr)
		sb.WriteString(row)
	}

	return sb.String(), nil
}


func formatMoney(cents int, currency, locale string) string {
	symbol := "$"
	if currency == "EUR" {
		symbol = "€"
	}

	isNegative := cents < 0
	if isNegative {
		cents = -cents
	}

	// Obtener la representacion de los miles y los centavos centesimales
	integerPart := cents / 100
	fractionalPart := cents % 100

	// Formatear la parte entera insertando los separadores de miles manualmente
	intStr := strconv.Itoa(integerPart)
	var thousandSeparator, decimalSeparator string
	if locale == "nl-NL" {
		thousandSeparator, decimalSeparator = ".", ","
	} else {
		thousandSeparator, decimalSeparator = ",", "."
	}

	var formattedInt strings.Builder
	for i, runeValue := range intStr {
		if i > 0 && (len(intStr)-i)%3 == 0 {
			formattedInt.WriteString(thousandSeparator)
		}
		formattedInt.WriteRune(runeValue)
	}

	// Reensamblar la cadena monetaria con su respectivo signo de puntuacion
	var moneyRaw string
	if locale == "nl-NL" {
		// Formato holandes: "€ 1.234,56 " o "€ -1.234,56 "
		sign := ""
		if isNegative {
			sign = "-"
		}
		moneyRaw = fmt.Sprintf("%s %s%s%s%02d ", symbol, sign, formattedInt.String(), decimalSeparator, fractionalPart)
	} else {
		// Formato americano: "$1,234.56 " o "($1,234.56)"
		if isNegative {
			moneyRaw = fmt.Sprintf("(%s%s%s%02d)", symbol, formattedInt.String(), decimalSeparator, fractionalPart)
		} else {
			moneyRaw = fmt.Sprintf("%s%s%s%02d ", symbol, formattedInt.String(), decimalSeparator, fractionalPart)
		}
	}

	return moneyRaw
}
