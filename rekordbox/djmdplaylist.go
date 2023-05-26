package rekordbox

import (
	"encoding/xml"
	"fmt"
	"math"
	"strconv"
)

type ConditionOperator string

const (
	ConditionOperatorContains    ConditionOperator = "8"
	ConditionOperatorNotContains ConditionOperator = "9"
	ConditionOperatorEndsWith    ConditionOperator = "11"
)

type LogicalOperator string

const (
	LogicalOperatorAllOf LogicalOperator = "1"
	LogicalOperatorAnyOf LogicalOperator = "2"
)

type PropertyName string

const (
	PropertyNameMyTag  PropertyName = "myTag"
	PropertyNameArtist PropertyName = "artist"
)

type DjmdPlaylistNodeCondition struct {
	PropertyName PropertyName      `xml:"PropertyName,attr"`
	Operator     ConditionOperator `xml:"Operator,attr"`
	ValueUnit    string            `xml:"ValueUnit,attr"`
	ValueLeft    string            `xml:"ValueLeft,attr"`
	ValueRight   string            `xml:"ValueRight,attr"`
}

// ParseValueLeft returns the correctly parsed MyTag ID for the left value
// Rekordbox stores values as int32 in the DB, so high values will overflow into a negative value
// this helper function corrects the value back to it's int64 form so it can be used
// for proper querying
func (c *DjmdPlaylistNodeCondition) ParseValueLeft() string {
	// only logic for MyTag is implemented currently
	// all the other conditions we can just return as is
	if c.PropertyName != PropertyNameMyTag {
		return c.ValueLeft
	}

	storedValue, err := strconv.ParseInt(c.ValueLeft, 10, 64)
	if err != nil {
		return c.ValueLeft
	}

	if storedValue >= 0 {
		return c.ValueLeft
	}

	originalValue := storedValue + int64(math.MaxInt32) + 1 // Add the max int32 value and 1
	originalValue += int64(math.MaxInt32) + 1               // Add the max int32 value and 1 again

	return fmt.Sprintf("%d", originalValue)
}

type DjmdPlaylistNode struct {
	Id              string                       `xml:"Id,attr"`
	LogicalOperator LogicalOperator              `xml:"LogicalOperator,attr"`
	AutomaticUpdate string                       `xml:"AutomaticUpdate,attr"`
	Conditions      []*DjmdPlaylistNodeCondition `xml:"CONDITION"`
}

// IsSmartlist returns a bool whether the current playlist is a smartlist or not
func (dp *DjmdPlaylist) IsSmartlist() bool {
	return dp.SmartList.Valid() == true
}

func (dp *DjmdPlaylist) SmartlistNode() *DjmdPlaylistNode {
	var n DjmdPlaylistNode
	xml.Unmarshal([]byte(dp.SmartList.String()), &n)

	return &n
}
