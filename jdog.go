package jdog

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrUnknownType = errors.New("unknown type")
	ErrNotFound    = errors.New("not found")
	ErrOutOfBounds = errors.New("array index out of bounds")
)

// Get returns a value at a specified path if it exists. If it doesn't exist it
// returns an error.
func Get(v interface{}, q string) (interface{}, error) {
	switch v.(type) {
	case []interface{}:
		return getArrVal(v.([]interface{}), q)
	case map[string]interface{}:
		return getMapVal(v.(map[string]interface{}), q)
	default:
		return nil, ErrUnknownType
	}
}

func getMapVal(m map[string]interface{}, q string) (interface{}, error) {
	selector, q := mapPart(q)
	v, ok := m[selector]
	if !ok {
		return nil, ErrNotFound
	}
	if len(q) == 0 {
		return v, nil
	}

	return Get(v, q)
}

func mapPart(q string) (string, string) {
	parts := strings.Split(q, ".")
	if i := strings.Index(parts[0], "["); i != -1 {
		p := parts[0][:i]
		parts[0] = parts[0][i:]
		return p, strings.Join(parts, ".")
	}

	return parts[0], strings.Join(parts[1:], ".")
}

func getArrVal(a []interface{}, q string) (interface{}, error) {
	i, q := arrPart(q)
	if i == -1 || i >= len(a) {
		return nil, ErrOutOfBounds
	}
	v := a[i]
	if len(q) == 0 {
		return v, nil
	}
	return Get(v, q)
}

var arrayRegex = regexp.MustCompile("^\\[([0-9]+)\\]")

func arrPart(q string) (int, string) {
	matches := arrayRegex.FindStringSubmatch(q)
	arrayRegex.FindStringSubmatch(q)
	if len(matches) != 2 {
		return -1, ""
	}

	i, err := strconv.Atoi(matches[1])
	if err != nil {
		return -1, ""
	}

	q = q[len(matches[0]):]
	if len(q) > 1 && q[0] == '.' {
		q = q[1:]
	}
	return i, q
}
