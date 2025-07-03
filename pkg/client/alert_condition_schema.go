package client

import "errors"

type AlertOperatorType string

const (
	AlertAggregationOperatorType             AlertOperatorType = "aggregationOperator"
	AlertBinaryOperatorType                  AlertOperatorType = "binaryOperator"
	AlertConstantValueType                   AlertOperatorType = "constantValue"
	AlertLogicalOperatorType                 AlertOperatorType = "logicalOperator"
	AlertAttributeType                       AlertOperatorType = "attributeField"
	AlertMetricFieldType                     AlertOperatorType = "metricField"
	AlertQueryFieldType                      AlertOperatorType = "queryField"
	AlertRelationshipOperatorType            AlertOperatorType = "relationshipOperator"
	AlertRelationshipAggregationOperatorType AlertOperatorType = "relationshipAggregationOperator"
	AlertScopeFieldType                      AlertOperatorType = "scopeField"
)

type AlertAggregationOperator string

const (
	AlertOperatorCount AlertAggregationOperator = "COUNT"
	AlertOperatorMin   AlertAggregationOperator = "MIN"
	AlertOperatorMax   AlertAggregationOperator = "MAX"
	AlertOperatorAvg   AlertAggregationOperator = "AVG"
	AlertOperatorSum   AlertAggregationOperator = "SUM"
	AlertOperatorLast  AlertAggregationOperator = "LAST"
)

type AlertBinaryOperator string

const (
	AlertOperatorEq AlertBinaryOperator = "="
	AlertOperatorNe AlertBinaryOperator = "!="
	AlertOperatorGt AlertBinaryOperator = ">"
	AlertOperatorLt AlertBinaryOperator = "<"
	AlertOperatorGe AlertBinaryOperator = ">="
	AlertOperatorLe AlertBinaryOperator = "<="
	AlertOperatorIn AlertBinaryOperator = "IN"
)

type AlertLogicalOperator string

const (
	AlertOperatorAnd AlertLogicalOperator = "AND"
	AlertOperatorOr  AlertLogicalOperator = "OR"
)

var (
	AlertOperators = map[AlertOperatorType][]string{
		AlertAggregationOperatorType: {
			string(AlertOperatorCount),
			string(AlertOperatorMin),
			string(AlertOperatorMax),
			string(AlertOperatorAvg),
			string(AlertOperatorSum),
			string(AlertOperatorLast),
		},
		AlertBinaryOperatorType: {
			string(AlertOperatorEq),
			string(AlertOperatorNe),
			string(AlertOperatorGt),
			string(AlertOperatorLt),
			string(AlertOperatorGe),
			string(AlertOperatorLe),
			string(AlertOperatorIn),
		},
		AlertLogicalOperatorType: {
			string(AlertOperatorAnd),
			string(AlertOperatorOr),
		},
	}

	FilterOperations = map[string]FilterOperation{
		string(AlertOperatorEq): FilterOperationEq,
		string(AlertOperatorNe): FilterOperationNe,
		string(AlertOperatorGt): FilterOperationGt,
		string(AlertOperatorGe): FilterOperationGe,
		string(AlertOperatorLt): FilterOperationLt,
		string(AlertOperatorLe): FilterOperationLe,
	}
)

func GetAlertConditionType(operator string) (string, error) {
	for operatorType, operatorArray := range AlertOperators {
		exists := sliceValueExists(operatorArray, operator)
		if exists {
			return string(operatorType), nil
		}
	}

	return "", errors.New("alert operation not supported")
}
