package poolAPIQueryUtils

import (
	"math/big"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	filtersProto "github.com/grandminingpool/pool-api-proto/generated/utils/filters"
)

const (
	IntegerRangeFilterSeparator  = ":"
	DateTimeRangeFilterSeparator = ".."
)

func ParseBigIntRangeFilter(value string) *filtersProto.BigIntRangeFilter {
	bigIntStrSplits := strings.Split(value, IntegerRangeFilterSeparator)

	if len(bigIntStrSplits) == 2 {
		minStr, maxStr := bigIntStrSplits[0], bigIntStrSplits[1]
		bigIntRangeFilter := &filtersProto.BigIntRangeFilter{}

		if minStr != "" {
			min, ok := new(big.Int).SetString(minStr, 10)
			if ok {
				bigIntRangeFilter.Min = min.Bytes()
			}
		}

		if maxStr != "" {
			max, ok := new(big.Int).SetString(maxStr, 10)
			if ok {
				bigIntRangeFilter.Max = max.Bytes()
			}
		}

		return bigIntRangeFilter
	}

	return nil
}

func ParseUInt32Filter(value string) *filtersProto.UInt32Filter {
	uint32StrSplits := strings.Split(value, IntegerRangeFilterSeparator)

	if len(uint32StrSplits) == 1 {
		equals64, err := strconv.ParseUint(value, 10, 32)
		if err == nil {
			equals32 := uint32(equals64)

			return &filtersProto.UInt32Filter{
				Operators: &filtersProto.UInt32Filter_Equals{
					Equals: equals32,
				},
			}
		}
	} else if len(uint32StrSplits) == 2 {
		minStr, maxStr := uint32StrSplits[0], uint32StrSplits[1]
		uint32Range := &filtersProto.UInt32Range{}

		if minStr != "" {
			min64, err := strconv.ParseUint(minStr, 10, 32)
			if err == nil {
				min32 := uint32(min64)
				uint32Range.Min = &min32
			}
		}

		if maxStr != "" {
			max64, err := strconv.ParseUint(maxStr, 10, 32)
			if err == nil {
				max32 := uint32(max64)
				uint32Range.Max = &max32
			}
		}

		return &filtersProto.UInt32Filter{
			Operators: &filtersProto.UInt32Filter_Range{
				Range: uint32Range,
			},
		}
	}

	return nil
}

func ParseUInt64RangeFilter(value string) *filtersProto.UInt64RangeFilter {
	uint64StrSplits := strings.Split(value, IntegerRangeFilterSeparator)

	if len(uint64StrSplits) == 2 {
		minStr, maxStr := uint64StrSplits[0], uint64StrSplits[1]
		uint64RangeFilter := &filtersProto.UInt64RangeFilter{}

		if minStr != "" {
			min, err := strconv.ParseUint(minStr, 10, 64)
			if err == nil {
				uint64RangeFilter.Min = &min
			}
		}

		if maxStr != "" {
			max, err := strconv.ParseUint(maxStr, 10, 64)
			if err == nil {
				uint64RangeFilter.Max = &max
			}
		}

		return uint64RangeFilter
	}

	return nil
}

func ParseDateTimeRangeFilter(value string) *filtersProto.DateTimeRangeFilter {
	dateTimeStrSplits := strings.Split(value, DateTimeRangeFilterSeparator)

	if len(dateTimeStrSplits) == 2 {
		startStr, endStr := dateTimeStrSplits[0], dateTimeStrSplits[1]
		dateTimeRangeFilter := &filtersProto.DateTimeRangeFilter{}

		if startStr != "" {
			start, err := time.Parse(time.RFC3339, startStr)
			if err == nil {
				dateTimeRangeFilter.Start = timestamppb.New(start)
			}
		}

		if endStr != "" {
			end, err := time.Parse(time.RFC3339, endStr)
			if err == nil {
				dateTimeRangeFilter.End = timestamppb.New(end)
			}
		}

		return dateTimeRangeFilter
	}

	return nil
}
