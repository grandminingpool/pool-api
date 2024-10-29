package poolAPIQueryUtils

import (
	"regexp"
	"strings"

	"go.uber.org/zap"

	sortsProto "github.com/grandminingpool/pool-api-proto/generated/utils/sorts"
)

var querySortItemRx *regexp.Regexp

func init() {
	re, err := regexp.Compile("([^:]+):([^:]+)")
	if err != nil {
		zap.L().Fatal("failed to compile query sort item regexp", zap.Error(err))
	}

	querySortItemRx = re
}

func ParseSortsItems(value string) map[string]*sortsProto.SortOrder {
	sortsMap := make(map[string]*sortsProto.SortOrder)
	sortItems := strings.Split(value, ",")

	for _, sortItem := range sortItems {
		matches := querySortItemRx.FindStringSubmatch(sortItem)

		if len(matches) == 3 {
			field := matches[1]
			direction := matches[2]
			sortDirection := sortsProto.SortDirection_DESC
			if direction == "asc" {
				sortDirection = sortsProto.SortDirection_ASC
			}

			sortsMap[field] = &sortsProto.SortOrder{
				Direction: sortDirection,
			}
		}
	}

	return sortsMap
}
