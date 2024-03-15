package Response

// Paginate function:  paginates the given CSV file records based on the page number, page size, and total records.
//
// records: a slice of string slices representing the records
// page: the page number
// pageSize: the number of records per page
// totalRecords: the total number of records in the file
// [][]string: a slice of string slices representing the paginated records
func Paginate(records [][]string, page int, pageSize int, totalRecords int) [][]string {
	start := (page - 1) * pageSize
	end := start + pageSize
	if end > totalRecords {
		end = totalRecords
	}
	return records[start+1 : end]
}

func Map(headers []string, records [][]string) []map[string]string {
	var data []map[string]string = []map[string]string{}
	for i := 0; i < len(records); i++ {
		recordMap := make(map[string]string)
		for j := 0; j < len(headers) && j < len(records[i]); j++ {
			if j < len(records[i]) {
				recordMap[headers[j]] = records[i][j]
			}
		}
		data = append(data, recordMap)
	}
	return data
}
