package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var positionMap = map[string]string{
	"1": "P", "2": "C", "3": "1B", "4": "2B", "5": "3B", "6": "SS",
	"7": "LF", "8": "CF", "9": "RF", "D": "DH", "H": "UT", // H = Utility/PH
}

func main() {
	inputRoot := "data/baseball-reference.com"
	outputDir := "data/prepared_data/"
	os.MkdirAll(outputDir, 0755)

	// List all team folders
	teamDirs, err := os.ReadDir(inputRoot)
	if err != nil {
		log.Fatalf("Failed to read input root: %v", err)
	}

	// List of table types you want to process
	tableTypes := []string{
		"standard_batting.csv",
		"value_batting.csv",
		"standard_pitching.csv",
		"value_pitching.csv",
		"standard_fielding.csv",
		"full_season_roster.csv", // Add this
		"40_man_roster.csv",      // And this (or current_40_roster.csv)
		// Add more as needed
	}

	for _, teamDir := range teamDirs {
		if !teamDir.IsDir() {
			continue
		}
		teamName := teamDir.Name()
		teamPath := filepath.Join(inputRoot, teamName)
		for _, table := range tableTypes {
			inputPath := filepath.Join(teamPath, table)
			if _, err := os.Stat(inputPath); err != nil {
				continue // skip missing files
			}
			// e.g., cubs_standard_batting_clean.csv
			shortTeam := strings.ToLower(strings.Fields(teamName)[2]) // "Cubs" or "White Sox"
			if shortTeam == "white" {                                 // handle "Chicago White Sox"
				shortTeam = "sox"
			}
			if shortTeam == "cubs" {
				shortTeam = "cubs"
			}
			tableBase := strings.TrimSuffix(table, ".csv")
			outputFile := filepath.Join(outputDir, fmt.Sprintf("%s_%s_clean.csv", shortTeam, tableBase))
			fmt.Printf("Processing %s -> %s\n", inputPath, outputFile)
			cleanAndWriteCSV(inputPath, outputFile)
		}
	}
}

func cleanAndWriteCSV(inputFile, outputFile string) {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Failed to open input: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, `"`, ` in`)
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to scan file: %v", err)
	}
	r := csv.NewReader(strings.NewReader(strings.Join(lines, "\n")))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV: %v", err)
	}

	// Find the real header row (skip stat family/sub-header rows)
	var realHeader []string
	headerIdx := -1
	for i, row := range records {
		for _, col := range row {
			colTrim := strings.TrimSpace(col)
			if colTrim == "Player" || colTrim == "Name" {
				realHeader = row
				headerIdx = i
				break
			}
		}
		if headerIdx != -1 {
			break
		}
	}
	if headerIdx == -1 {
		log.Fatalf("Could not find real header row in %s", inputFile)
	}
	header := realHeader

	// Detect per-position columns
	positionColumns := []string{"P", "C", "1B", "2B", "3B", "SS", "LF", "CF", "RF", "OF", "DH", "UT"}
	positionColIdx := map[string]int{}
	posIdx := -1
	for i, col := range header {
		for _, pos := range positionColumns {
			if col == pos {
				positionColIdx[pos] = i
			}
		}
		if strings.ToLower(col) == "pos" {
			posIdx = i
		}
	}

	// Detect handedness columns
	batHandIdx, playerIdx := -1, -1
	for i, col := range header {
		switch strings.ToLower(col) {
		case "b", "bats":
			batHandIdx = i
		case "player", "name":
			playerIdx = i
		}
	}

	// Build new header
	var newHeader []string
	for i, col := range header {
		if i == playerIdx {
			newHeader = append(newHeader, col, "Handedness")
		} else {
			newHeader = append(newHeader, col)
		}
	}
	// Add EligiblePositions, PositionFlexibility, and boolean columns if per-position columns exist or Pos column exists
	if len(positionColIdx) > 0 || posIdx != -1 {
		newHeader = append(newHeader, "EligiblePositions", "PositionFlexibility")
		for _, pos := range positionColumns {
			newHeader = append(newHeader, "Is"+pos)
		}
	}

	w, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create output: %v", err)
	}
	defer w.Close()
	writer := csv.NewWriter(w)
	writer.Write(newHeader)

	// Process rows
	for _, row := range records[headerIdx+1:] {
		// Skip stat family/sub-header rows and summary rows
		if len(row) > 1 && (strings.TrimSpace(row[1]) == "Standard" ||
			strings.Contains(strings.ToLower(row[1]), "shift") ||
			strings.Contains(strings.ToLower(row[1]), "positioning")) {
			continue
		}
		// Skip empty rows
		if len(row) == 0 || (len(row) > 1 && strings.TrimSpace(row[1]) == "") {
			continue
		}
		// Skip team totals row (handled separately)
		if len(row) > 1 && strings.EqualFold(strings.TrimSpace(row[1]), "Team Totals") {
			continue
		}
		var newRow []string
		for i, val := range row {
			if i == playerIdx {
				// Stat tables: parse handedness from Player/Name
				// Roster tables: use B column if present
				hand := ""
				if batHandIdx != -1 && batHandIdx < len(row) && row[batHandIdx] != "" {
					hand = row[batHandIdx]
					name, _ := parsePlayerHandedness(val)
					newRow = append(newRow, name, hand)
				} else {
					name, handParsed := parsePlayerHandedness(val)
					newRow = append(newRow, name, handParsed)
				}
			} else {
				newRow = append(newRow, val)
			}
		}
		// EligiblePositions, PositionFlexibility, and boolean columns
		eligiblePositions := []string{}
		posSet := make(map[string]bool)
		// Use per-position columns if present, otherwise parse Pos column
		if len(positionColIdx) > 0 {
			for _, pos := range positionColumns {
				idx, ok := positionColIdx[pos]
				if ok && idx < len(row) && row[idx] != "" && row[idx] != "0" {
					eligiblePositions = append(eligiblePositions, pos)
					posSet[pos] = true
				}
			}
		} else if posIdx != -1 && posIdx < len(row) {
			// Parse Pos column (e.g., "*3/H" or "8/D")
			posList := parsePositions(row[posIdx])
			for _, pos := range posList {
				eligiblePositions = append(eligiblePositions, pos)
				posSet[pos] = true
			}
		}
		if len(positionColIdx) > 0 || posIdx != -1 {
			newRow = append(newRow, strings.Join(eligiblePositions, ","), fmt.Sprintf("%d", len(eligiblePositions)))
			for _, pos := range positionColumns {
				if posSet[pos] {
					newRow = append(newRow, "1")
				} else {
					newRow = append(newRow, "0")
				}
			}
		}
		writer.Write(newRow)
	}
	writer.Flush()
	fmt.Printf("Wrote cleaned CSV: %s\n", outputFile)

	// After processing all rows, extract team totals if present
	var teamTotalsRow []string
	for _, row := range records[headerIdx+1:] {
		if len(row) > 1 && strings.EqualFold(strings.TrimSpace(row[1]), "Team Totals") {
			// Build newRow for team totals using the same logic as above
			var newRow []string
			for i, val := range row {
				if i == playerIdx {
					hand := ""
					if batHandIdx != -1 && batHandIdx < len(row) && row[batHandIdx] != "" {
						hand = row[batHandIdx]
						name, _ := parsePlayerHandedness(val)
						newRow = append(newRow, name, hand)
					} else {
						name, handParsed := parsePlayerHandedness(val)
						newRow = append(newRow, name, handParsed)
					}
				} else {
					newRow = append(newRow, val)
				}
			}
			eligiblePositions := []string{}
			posSet := make(map[string]bool)
			if len(positionColIdx) > 0 {
				for _, pos := range positionColumns {
					idx, ok := positionColIdx[pos]
					if ok && idx < len(row) && row[idx] != "" && row[idx] != "0" {
						eligiblePositions = append(eligiblePositions, pos)
						posSet[pos] = true
					}
				}
			} else if posIdx != -1 && posIdx < len(row) {
				posList := parsePositions(row[posIdx])
				for _, pos := range posList {
					eligiblePositions = append(eligiblePositions, pos)
					posSet[pos] = true
				}
			}
			if len(positionColIdx) > 0 || posIdx != -1 {
				newRow = append(newRow, strings.Join(eligiblePositions, ","), fmt.Sprintf("%d", len(eligiblePositions)))
				for _, pos := range positionColumns {
					if posSet[pos] {
						newRow = append(newRow, "1")
					} else {
						newRow = append(newRow, "0")
					}
				}
			}
			teamTotalsRow = newRow
			break
		}
	}
	if len(teamTotalsRow) > 0 && len(realHeader) > 0 {
		totalsFile := strings.Replace(outputFile, "_clean.csv", "_team_totals_clean.csv", 1)
		f, err := os.Create(totalsFile)
		if err != nil {
			log.Printf("Failed to create team totals file %s: %v", totalsFile, err)
		} else {
			writer := csv.NewWriter(f)
			writer.Write(newHeader) // <-- this is correct!
			writer.Write(teamTotalsRow)
			writer.Flush()
			f.Close()
			fmt.Printf("Wrote team totals CSV: %s\n", totalsFile)
		}
	}
}

func parsePlayerHandedness(player string) (string, string) {
	player = strings.TrimSpace(player)
	if strings.HasSuffix(player, "*") {
		return strings.TrimSuffix(player, "*"), "L"
	}
	if strings.HasSuffix(player, "#") {
		return strings.TrimSuffix(player, "#"), "S"
	}
	return player, "R"
}

func parsePositions(posField string) []string {
	// Remove leading * or /, split on /, map to standard codes
	posField = strings.TrimLeft(posField, "*/")
	parts := regexp.MustCompile(`[\/]`).Split(posField, -1)
	var positions []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		// Map numbers/letters to standard positions
		for _, ch := range p {
			if pos, ok := positionMap[string(ch)]; ok {
				positions = append(positions, pos)
			}
		}
	}
	return uniqueStrings(positions)
}

func uniqueStrings(input []string) []string {
	seen := make(map[string]struct{})
	var out []string
	for _, v := range input {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

func isLikelyPlayerName(s string) bool {
	// crude check: not a stat family, not a number, not empty
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}
	lower := strings.ToLower(s)
	if lower == "standard" || lower == "total zone" || lower == "drs" || lower == "range factor" || lower == "baserunners" {
		return false
	}
	// Not a number
	if _, err := strconv.Atoi(s); err == nil {
		return false
	}
	return true
}
