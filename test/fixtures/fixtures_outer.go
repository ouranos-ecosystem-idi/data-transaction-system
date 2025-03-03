package fixtures

import "fmt"

func GetPartsStructure_AllItem() string {
	return `{
		"parent": {
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsItem": "B01",
			"supportPartsItem": "A000001",
			"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
			"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"amountUnitName": "kilogram",
			"endFlag": false,
			"partsLabelName": "PartsB",
			"partsAddInfo1": "Ver3.0",
			"partsAddInfo2": "2024-12-01-2024-12-31",
			"partsAddInfo3": "任意の情報が入ります"
		},
		"children": [
			{
				"partsStructureId": "2680ed32-19a3-435b-a094-23ff43aaa611_1c2f37f5-25b9-dea5-346a-7b88035f2553",
				"traceId": "1c2f37f5-25b9-dea5-346a-7b88035f2553",
				"partsItem": "B01001",
				"supportPartsItem": "B001",
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountUnitName": "kilogram",
				"endFlag": false,
				"amount": 2.1,
				"partsLabelName": "PartsB",
				"partsAddInfo1": "Ver3.0",
				"partsAddInfo2": "2024-12-01-2024-12-31",
				"partsAddInfo3": "任意の情報が入ります"
			}
		]
	}`
}

func GetPartsStructure_RequireItemOnly() string {
	return `{
		"parent": {
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsItem": "B01",
			"supportPartsItem": null,
			"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
			"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"amountUnitName": null,
			"endFlag": false,
			"partsLabelName": null,
			"partsAddInfo1": null,
			"partsAddInfo2": null,
			"partsAddInfo3": null
		},
		"children": [
			{
				"partsStructureId": "2680ed32-19a3-435b-a094-23ff43aaa611_1c2f37f5-25b9-dea5-346a-7b88035f2553",
				"traceId": "1c2f37f5-25b9-dea5-346a-7b88035f2553",
				"partsItem": "B01001",
				"supportPartsItem": null,
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountUnitName": null,
				"endFlag": false,
				"amount": null,
				"partsLabelName": null,
				"partsAddInfo1": null,
				"partsAddInfo2": null,
				"partsAddInfo3": null
			}
		]
	}`
}

func GetPartsStructure_RequireItemOnlyWithUndefined() string {
	return `{
		"parent": {
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsItem": "B01",
			"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
			"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"endFlag": false
		},
		"children": [
			{
				"partsStructureId": "2680ed32-19a3-435b-a094-23ff43aaa611_1c2f37f5-25b9-dea5-346a-7b88035f2553",
				"traceId": "1c2f37f5-25b9-dea5-346a-7b88035f2553",
				"partsItem": "B01001",
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"endFlag": false
			}
		]
	}`
}

func GetPartsStructure_NoComponent() string {
	return `{
		"parent": {
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsItem": "B01",
			"supportPartsItem": null,
			"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
			"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"amountUnitName": null,
			"endFlag": false,
			"partsLabelName": null,
			"partsAddInfo1": null,
			"partsAddInfo2": null,
			"partsAddInfo3": null
		},
		"children": []
	}`
}

func GetPartsStructure_NoData() string {
	return `{
		"parent": null,
		"children": []
	}`
}

func GetPartsStructure_InvalidTypeError() string {
	return `{
		"parent": {
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsItem": "B01",
			"supportPartsItem": "A000001",
			"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
			"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
			"amountUnitName": "kg",
			"endFlag": false,
			"partsLabelName": "PartsB",
			"partsAddInfo1": "Ver3.0",
			"partsAddInfo2": "2024-12-01-2024-12-31",
			"partsAddInfo3": "任意の情報が入ります"
		},
		"children": [
			{
				"partsStructureId": "2680ed32-19a3-435b-a094-23ff43aaa611_1c2f37f5-25b9-dea5-346a-7b88035f2553",
				"traceId": "1c2f37f5-25b9-dea5-346a-7b88035f2553",
				"partsItem": "B01001",
				"supportPartsItem": "B001",
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountUnitName": "kg",
				"endFlag": false,
				"amount": 2.1,
				"partsLabelName": "PartsB",
				"partsAddInfo1": "Ver3.0",
				"partsAddInfo2": "2024-12-01-2024-12-31",
				"partsAddInfo3": "任意の情報が入ります"
			}
		]
	}`
}

func PutPartsStructure() string {
	return `{
		"parent": {
			"traceId": "d17833fe-22b7-4a4a-b097-bc3f2150c9a6",
			"partsItem": "PartsA-002123",
			"supportPartsItem": "modelA",
			"partsLabelName": "PartsB",
			"partsAddInfo1": "Ver3.0",
			"partsAddInfo2": "2024-12-01-2024-12-31",
			"partsAddInfo3": "任意の情報が入ります"
		},
		"children": [
			{
				"partsStructureId": "d17833fe-22b7-4a4a-b097-bc3f2150c9a6_06c9b015-4225-ba30-1ed3-6faf02cb3fe6",
				"traceId": "06c9b015-4225-ba30-1ed3-6faf02cb3fe6",
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"partsItem": "PartsA-002123",
				"supportPartsItem": "modelA",
				"partsLabelName": "PartsB",
				"partsAddInfo1": "Ver3.0",
				"partsAddInfo2": "2024-12-01-2024-12-31",
				"partsAddInfo3": "任意の情報が入ります"
			}
		]
	}`
}

func GetParts_AllItem(inputTraceId *string) string {
	traceId := "2680ed32-19a3-435b-a094-23ff43aaa611"
	if inputTraceId != nil {
		traceId = *inputTraceId
	}
	return fmt.Sprintf(`{
		"parts": [
			{
				"traceId": "%s",
				"partsItem": "B01",
				"supportPartsItem": "A000001",
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountUnitName": "kilogram",
				"endFlag": false,
				"parentFlag": true,
				"partsLabelName": "PartsB",
				"partsAddInfo1": "Ver3.0",
				"partsAddInfo2": "2024-12-01-2024-12-31",
				"partsAddInfo3": "任意の情報が入ります"

			}
		],
		"next": "2680ed32-19a3-435b-a094-23ff43aaa612"
	}`, traceId)
}

func GetParts_TradeRequestParts(inputTraceId *string) string {
	traceId := "2680ed32-19a3-435b-a094-23ff43aaa611"
	if inputTraceId != nil {
		traceId = *inputTraceId
	}
	return fmt.Sprintf(`{
		"parts": [
			{
				"traceId": "%s",
				"partsItem": "B01",
				"supportPartsItem": "A000001",
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountUnitName": "kilogram",
				"endFlag": false,
				"parentFlag": false
				"partsLabelName": "PartsB",
				"partsAddInfo1": "Ver3.0",
				"partsAddInfo2": "2024-12-01-2024-12-31",
				"partsAddInfo3": "任意の情報が入ります"
			}
		],
		"next": "2680ed32-19a3-435b-a094-23ff43aaa612"
	}`, traceId)
}

func GetParts_RequireItemOnly() string {
	return `{
		"parts": [
			{
				"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
				"partsItem": "B01",
				"supportPartsItem": null,
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountUnitName": null,
				"endFlag": false,
				"parentFlag": true,
				"partsLabelName": null,
				"partsAddInfo1": null,
				"partsAddInfo2": null,
				"partsAddInfo3": null
			}
		],
		"next": null
	}`
}

func GetParts_RequireItemOnlyWithUndefined() string {
	return `{
		"parts": [
			{
				"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
				"partsItem": "B01",
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"endFlag": false,
				"parentFlag": true
			}
		]
	}`
}

func GetParts_NoData() string {
	return `{
		"parts": [],
		"next": null
	}`
}

func GetParts_InvalidTypeError() string {
	return `{
		"parts": [
			{
				"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
				"partsItem": "B01",
				"supportPartsItem": "A000001",
				"plantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
				"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
				"amountUnitName": "kg",
				"endFlag": false,
				"parentFlag": false,
				"partsLabelName": "PartsB",
				"partsAddInfo1": "Ver3.0",
				"partsAddInfo2": "2024-12-01-2024-12-31",
				"partsAddInfo3": "任意の情報が入ります"
			}
		],
		"next": "2680ed32-19a3-435b-a094-23ff43aaa612"
	}`
}

func GetTradeRequests_AllItem() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedToOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": "A01のCFP値を回答ください",
					"replyMessage": "A01のCFP値を回答しました",
					"responseDueDate": "2024-12-31",
					"completedCount": 0,
					"completedCountModifiedAt": "2024-05-23T11:22:33Z"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"upstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b"
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamAmountUnitName": "kilogram"
					},
					"tradesCount": 0,
					"tradesCountModifiedAt": "2024-05-24T22:33:44Z"
				},
				"response": {
					"responseId": "b26e3bd3-7443-4f23-8cce-2056052f0452",
					"responseType": "CFP",
					"responsedAt": "2024-02-16T17:46:21Z",
					"responsePreProcessingEmissions": 0.1,
					"responseMainProductionEmissions": 0.4,
					"emissionsUnitName": "kgCO2e/kilogram",
					"cfpCertificationFileInfo": [
						{
							"fileId": "fe517a2b-2af8-48ff-b1ed-88fc50f4414f",
							"fileName": "B01_CFP.pdf"
						}
					],
					"responseDqr": {
						"preProcessingTeR": 3.1,
                        "preProcessingGeR": 3.2,
						"preProcessingTiR": 3.3,
						"mainProductionTeR": 3.4,
						"mainProductionGeR": 3.5,
						"mainProductionTiR": 3.6
					}
				}
			}
		],
		"next": "026ad6a0-a689-4b8c-8a14-7304b817096d"
	}`
}

func GetTradeRequests_AllItem_WithNull() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedToOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": "A01のCFP値を回答ください",
					"replyMessage": "A01のCFP値を回答しました",
					"responseDueDate": null,
					"completedCount": null,
					"completedCountModifiedAt": null
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"upstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b"
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamAmountUnitName": "kilogram"
					},
					"tradesCount": null,
					"tradesCountModifiedAt": null
				},
				"response": {
					"responseId": "b26e3bd3-7443-4f23-8cce-2056052f0452",
					"responseType": "CFP",
					"responsedAt": "2024-02-16T17:46:21Z",
					"responsePreProcessingEmissions": 0.1,
					"responseMainProductionEmissions": 0.4,
					"emissionsUnitName": "kgCO2e/kilogram",
					"cfpCertificationFileInfo": [
						{
							"fileId": "fe517a2b-2af8-48ff-b1ed-88fc50f4414f",
							"fileName": "B01_CFP.pdf"
						}
					],
					"responseDqr": {
						"preProcessingTeR": 3.1,
                        "preProcessingGeR": 3.2,
						"preProcessingTiR": 3.3,
						"mainProductionTeR": 3.4,
						"mainProductionGeR": 3.5,
						"mainProductionTiR": 3.6
					}
				}
			}
		],
		"next": "026ad6a0-a689-4b8c-8a14-7304b817096d"
	}`
}

func GetTradeRequests_AllItem_WithUndefined() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedToOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": "A01のCFP値を回答ください",
					"replyMessage": "A01のCFP値を回答しました"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"upstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b"
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamAmountUnitName": "kilogram"
					}
				},
				"response": {
					"responseId": "b26e3bd3-7443-4f23-8cce-2056052f0452",
					"responseType": "CFP",
					"responsedAt": "2024-02-16T17:46:21Z",
					"responsePreProcessingEmissions": 0.1,
					"responseMainProductionEmissions": 0.4,
					"emissionsUnitName": "kgCO2e/kilogram",
					"cfpCertificationFileInfo": [
						{
							"fileId": "fe517a2b-2af8-48ff-b1ed-88fc50f4414f",
							"fileName": "B01_CFP.pdf"
						}
					],
					"responseDqr": {
						"preProcessingTeR": 3.1,
                        "preProcessingGeR": 3.2,
						"preProcessingTiR": 3.3,
						"mainProductionTeR": 3.4,
						"mainProductionGeR": 3.5,
						"mainProductionTiR": 3.6
					}
				}
			}
		],
		"next": "026ad6a0-a689-4b8c-8a14-7304b817096d"
	}`
}

func GetTradeRequests_AllItem_NoNext() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedToOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": "A01のCFP値を回答ください",
					"replyMessage": "A01のCFP値を回答しました",
					"responseDueDate": "2024-12-31",
					"completedCount": 0,
					"completedCountModifiedAt": "2024-05-23T11:22:33Z"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"upstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b"
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamAmountUnitName": "kilogram"
					},
					"tradesCount": 0,
					"tradesCountModifiedAt": "2024-05-24T22:33:44Z"
				},
				"response": {
					"responseId": "b26e3bd3-7443-4f23-8cce-2056052f0452",
					"requestType": "CFP",
					"responsedAt": "2024-02-16T17:46:21Z",
					"responsePreProcessingEmissions": 0.1,
					"responseMainProductionEmissions": 0.4,
					"emissionsUnitName": "kgCO2e/kilogram",
					"cfpCertificationFileInfo": [
						{
							"fileId": "fe517a2b-2af8-48ff-b1ed-88fc50f4414f",
							"fileName": "B01_CFP.pdf"
						}
					],
					"responseDqr": {
						"preProcessingTeR": 3.1,
                        "preProcessingGeR": 3.2,
						"preProcessingTiR": 3.3,
						"mainProductionTeR": 3.4,
						"mainProductionGeR": 3.5,
						"mainProductionTiR": 3.6
					}
				}
			}
		],
		"next": null
	}`
}

func GetTradeRequests_RequireItemOnlyAnswered() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedToOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": null,
					"replyMessage": null,
					"responseDueDate": "2024-12-31",
					"completedCount": 0,
					"completedCountModifiedAt": "2024-05-23T11:22:33Z"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"upstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b"
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamAmountUnitName": "kilogram"
					},
					"tradesCount": 0,
					"tradesCountModifiedAt": "2024-05-24T22:33:44Z"
				},
				"response": {
					"responseId": "b26e3bd3-7443-4f23-8cce-2056052f0452",
					"requestType": "CFP",
					"responsedAt": "2024-02-16T17:46:21Z",
					"responsePreProcessingEmissions": null,
					"responseMainProductionEmissions": null,
					"emissionsUnitName": "kgCO2e/kilogram",
					"cfpCertificationFileInfo": [],
					"responseDqr": {
						"preProcessingTeR": null,
                        "preProcessingGeR": null,
						"preProcessingTiR": null,
						"mainProductionTeR": null,
						"mainProductionGeR": null,
						"mainProductionTiR": null
					}
				}
			}
		],
		"next": null
	}`
}

func GetTradeRequests_RequireItemOnlyAnsweredWithUndefined() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedToOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"requestedAt": "2024-02-14T15:25:35Z",
					"responseDueDate": "2024-12-31",
					"completedCount": 0,
					"completedCountModifiedAt": "2024-05-23T11:22:33Z"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"upstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b"
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamAmountUnitName": "kilogram"
					},
					"tradesCount": 0,
					"tradesCountModifiedAt": "2024-05-24T22:33:44Z"
				},
				"response": {
					"responseId": "b26e3bd3-7443-4f23-8cce-2056052f0452",
					"requestType": "CFP",
					"responsedAt": "2024-02-16T17:46:21Z",
					"emissionsUnitName": "kgCO2e/kilogram",
					"cfpCertificationFileInfo": [],
					"responseDqr": {}
				}
			}
		]
	}`
}

func GetTradeRequests_RequireItemOnlyAnswering() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedToOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": null,
					"replyMessage": null,
					"responseDueDate": "2024-12-31",
					"completedCount": 0,
					"completedCountModifiedAt": "2024-05-23T11:22:33Z"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"upstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": null
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamAmountUnitName": "kilogram"
					},
					"tradesCount": 0,
					"tradesCountModifiedAt": "2024-05-24T22:33:44Z"
				},
				"response": null
			}
		],
		"next": null
	}`
}

func GetTradeRequests_RequireItemOnlyAnsweringWithUndefined() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedToOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"requestedAt": "2024-02-14T15:25:35Z",
					"responseDueDate": "2024-12-31",
					"completedCount": 0,
					"completedCountModifiedAt": "2024-05-23T11:22:33Z"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"upstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5"
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamAmountUnitName": "kilogram"
					},
					"tradesCount": 0,
					"tradesCountModifiedAt": "2024-05-24T22:33:44Z"
				}
			}
		]
	}`
}

func GetTradeRequests_NoData() string {
	return `{
		"tradeRequests": [],
		"next": null
	}`
}

func PutTradeRequests() string {
	return `[
		{
			"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
			"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
			"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5"
		}
	]`
}

func GetTradeRequestsReceived_AllItem() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedFromOperatorId": "b1234567-1234-1234-1234-123456789012",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": "A01のCFP値を回答ください",
					"replyMessage": "A01のCFP値を回答しました",
					"responseDueDate": "2024-12-31",
					"completedCount": 0,
					"completedCountModifiedAt": "2024-05-23T11:22:33Z"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"downstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b"
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamPartsItem": "B01",
						"downstreamSupportPartsItem": "B0100",
						"downstreamPlantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
						"downstreamAmountUnitName": "kilogram",
						"downstreamPartsLabelName": "PartsB",
						"downstreamPartsAddInfo1": "Ver3.0",
						"downstreamPartsAddInfo2": "2024-12-01-2024-12-31",
						"downstreamPartsAddInfo3": "任意の情報が入ります"
					},
					"tradesCount": 0,
					"tradesCountModifiedAt": "2024-05-24T22:33:44Z"
				}
			}
		],
		"next": "026ad6a0-a689-4b8c-8a14-7304b817096d"
	}`
}

func GetTradeRequestsReceived_AllItem_WithNull() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedFromOperatorId": "b1234567-1234-1234-1234-123456789012",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": "A01のCFP値を回答ください",
					"replyMessage": "A01のCFP値を回答しました",
					"responseDueDate": null,
					"completedCount": null,
					"completedCountModifiedAt": null
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"downstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b"
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamPartsItem": "B01",
						"downstreamSupportPartsItem": "B0100",
						"downstreamPlantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
						"downstreamAmountUnitName": "kilogram",
						"downstreamPartsLabelName": null,
						"downstreamPartsAddInfo1": null,
						"downstreamPartsAddInfo2": null,
						"downstreamPartsAddInfo3": null
					},
					"tradesCount": null,
					"tradesCountModifiedAt": null
				}
			}
		],
		"next": "026ad6a0-a689-4b8c-8a14-7304b817096d"
	}`
}

func GetTradeRequestsReceived_AllItem_WithUndefined() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedFromOperatorId": "b1234567-1234-1234-1234-123456789012",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": "A01のCFP値を回答ください",
					"replyMessage": "A01のCFP値を回答しました"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"downstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b"
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamPartsItem": "B01",
						"downstreamSupportPartsItem": "B0100",
						"downstreamPlantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
						"downstreamAmountUnitName": "kilogram"
					}
				}
			}
		],
		"next": "026ad6a0-a689-4b8c-8a14-7304b817096d"
	}`
}

func GetTradeRequestsReceived_AllItem_NoNext() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedFromOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": "A01のCFP値を回答ください",
					"replyMessage": "A01のCFP値を回答しました",
					"responseDueDate": "2024-12-31",
					"completedCount": 0,
					"completedCountModifiedAt": "2024-05-23T11:22:33Z"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"downstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b"
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamPartsItem": "B01",
						"downstreamSupportPartsItem": "B0100",
						"downstreamPlantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
						"downstreamAmountUnitName": "kilogram",
						"downstreamPartsLabelName": "PartsB",
						"downstreamPartsAddInfo1": "Ver3.0",
						"downstreamPartsAddInfo2": "2024-12-01-2024-12-31",
						"downstreamPartsAddInfo3": "任意の情報が入ります"
					},
					"tradesCount": 0,
					"tradesCountModifiedAt": "2024-05-24T22:33:44Z"
				}
			}
		],
		"next": null
	}`
}

func GetTradeRequestsReceived_AllItem_MaxLength() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedFromOperatorId": "b1234567-1234-1234-1234-123456789012",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": "１０００文字ああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああ",
					"replyMessage": "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
					"responseDueDate": "2024-12-31",
					"completedCount": 0,
					"completedCountModifiedAt": "2024-05-23T11:22:33Z"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"downstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b"
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamPartsItem": "５０文字ああああああああああああああああああああああああああああああああああああああああああああああ",
						"downstreamSupportPartsItem": "５０文字ああああああああああああああああああああああああああああああああああああああああああああああ",
						"downstreamPlantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
						"downstreamAmountUnitName": "kilogram",
						"downstreamPartsLabelName": "５０文字ああああああああああああああああああああああああああああああああああああああああああああああ",
						"downstreamPartsAddInfo1": "５０文字ああああああああああああああああああああああああああああああああああああああああああああああ",
						"downstreamPartsAddInfo2": "５０文字ああああああああああああああああああああああああああああああああああああああああああああああ",
						"downstreamPartsAddInfo3": "５０文字ああああああああああああああああああああああああああああああああああああああああああああああ"
					},
					"tradesCount": 0,
					"tradesCountModifiedAt": "2024-05-24T22:33:44Z"
				}
			}
		],
		"next": "026ad6a0-a689-4b8c-8a14-7304b817096d"
	}`
}

func GetTradeRequestsReceived_RequireItemOnly() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedFromOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": null,
					"replyMessage": null,
					"responseDueDate": "2024-12-31",
					"completedCount": 0,
					"completedCountModifiedAt": "2024-05-23T11:22:33Z"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"downstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": null
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamPartsItem": "B01",
						"downstreamSupportPartsItem": null,
						"downstreamPlantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
						"downstreamAmountUnitName": "kilogram",
						"downstreamPartsLabelName": null,
						"downstreamPartsAddInfo1": null,
						"downstreamPartsAddInfo2": null,
						"downstreamPartsAddInfo3": null
					},
					"tradesCount": 0,
					"tradesCountModifiedAt": "2024-05-24T22:33:44Z"
				}
			}
		],
		"next": null
	}`
}

func GetTradeRequestsReceived_RequireItemOnlyWithUndefined() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedFromOperatorId": "b1234567-1234-1234-1234-123456789012",
					"requestedAt": "2024-02-14T15:25:35Z",
					"responseDueDate": "2024-12-31",
					"completedCount": 0,
					"completedCountModifiedAt": "2024-05-23T11:22:33Z"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"downstreamOperatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5"
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamPartsItem": "B01",
						"downstreamPlantId": "eedf264e-cace-4414-8bd3-e10ce1c090e0",
						"downstreamAmountUnitName": "kilogram"
					},
					"tradesCount": 0,
					"tradesCountModifiedAt": "2024-05-24T22:33:44Z"
				}
			}
		]
	}`
}

func GetTradeRequestsReceived_NoData() string {
	return `{
		"tradeRequests": [],
		"next": null
	}`
}

func PostTrades() string {
	return `{
		"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092"
	}`
}

func PutPostTradeRequestsCancelResponse() string {
	return `[
		{
			"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
			"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092"
		}
	]`
}

func PutPostTradeRequestsRejectResponse() string {
	return `[
		{
			"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
			"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092"
		}
	]`
}

func GetCfp_AllItem() string {
	return `[
		{
			"cfp": {
				"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
				"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
				"preProcessingOwnOriginatedEmissions": 0.5,
				"mainProductionOwnOriginatedEmissions": 0.6,
				"preProcessingSupplierOriginatedEmissions": 1.1,
				"mainProductionSupplierOriginatedEmissions": 1.2,
				"emissionsUnitName": "kgCO2e/kilogram",
				"cfpComment": "部品B01001のCFPのCFP情報コメント",
				"dqr": {
					"preProcessingTeR": 3.1,
					"preProcessingGeR": 3.2,
					"preProcessingTiR": 3.3,
					"mainProductionTeR": 3.4,
					"mainProductionGeR": 3.5,
					"mainProductionTiR": 3.6
				},
				"parentFlag": true
			},
			"totalCfp": {
				"totalPreProcessingOwnOriginatedEmissions": 1.5,
				"totalMainProductionOwnOriginatedEmissions": 1.6,
				"totalPreProcessingSupplierOriginatedEmissions": 2.1,
				"totalMainProductionSupplierOriginatedEmissions": 2.2,
				"emissionsUnitName": "kgCO2e/kilogram",
				"totalDqr": {
					"preProcessingTeR": 4.1,
					"preProcessingGeR": 4.2,
					"preProcessingTiR": 4.3,
					"mainProductionTeR": 4.4,
					"mainProductionGeR": 4.5,
					"mainProductionTiR": 4.6
				}
			}
		}
	]`
}

func GetCfp_RequireItemOnly() string {
	return `[
		{
			"cfp": {
				"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
				"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
				"preProcessingOwnOriginatedEmissions": 0.5,
				"mainProductionOwnOriginatedEmissions": 0.6,
				"preProcessingSupplierOriginatedEmissions": 1.1,
				"mainProductionSupplierOriginatedEmissions": 1.2,
				"emissionsUnitName": "kgCO2e/kilogram",
				"cfpComment": null,
				"dqr": {
					"preProcessingTeR": 3.1,
					"preProcessingGeR": 3.2,
					"preProcessingTiR": 3.3,
					"mainProductionTeR": 3.4,
					"mainProductionGeR": 3.5,
					"mainProductionTiR": 3.6
				},
				"parentFlag": true
			},
			"totalCfp": {
				"totalPreProcessingOwnOriginatedEmissions": 1.5,
				"totalMainProductionOwnOriginatedEmissions": 1.6,
				"totalPreProcessingSupplierOriginatedEmissions": 2.1,
				"totalMainProductionSupplierOriginatedEmissions": 2.2,
				"emissionsUnitName": "kgCO2e/kilogram",
				"totalDqr": {
					"preProcessingTeR": null,
					"preProcessingGeR": null,
					"preProcessingTiR": null,
					"mainProductionTeR": null,
					"mainProductionGeR": null,
					"mainProductionTiR": null
				}
			}
		}
	]`
}

func GetCfp_RequireItemOnlyWithUndefined() string {
	return `[
		{
			"cfp": {
				"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31",
				"traceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b",
				"preProcessingOwnOriginatedEmissions": 0.5,
				"mainProductionOwnOriginatedEmissions": 0.6,
				"preProcessingSupplierOriginatedEmissions": 1.1,
				"mainProductionSupplierOriginatedEmissions": 1.2,
				"emissionsUnitName": "kgCO2e/kilogram",
				"dqr": {
					"preProcessingTeR": 3.1,
					"preProcessingGeR": 3.2,
					"preProcessingTiR": 3.3,
					"mainProductionTeR": 3.4,
					"mainProductionGeR": 3.5,
					"mainProductionTiR": 3.6
				},
				"parentFlag": true
			},
			"totalCfp": {
				"totalPreProcessingOwnOriginatedEmissions": 1.5,
				"totalMainProductionOwnOriginatedEmissions": 1.6,
				"totalPreProcessingSupplierOriginatedEmissions": 2.1,
				"totalMainProductionSupplierOriginatedEmissions": 2.2,
				"emissionsUnitName": "kgCO2e/kilogram",
				"totalDqr": {}
			}
		}
	]`
}

func GetCfp_NoData() string {
	return `[]`
}

func PutCfp_AllItem(traceId string) string {
	return fmt.Sprintf(`[
		{
			"traceId": "%s",
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb31"
		}
	]`, traceId)
}

func PutCfp_AllItem_InvalidCfp(traceId string) string {
	return fmt.Sprintf(`[
		{
			"traceId": "%s",
			"cfpId": "892262ab-6795-4a97-bf25-d92c512ebb3"
		}
	]`, traceId)
}

func DeleteParts_AllItem(traceId string) string {
	return fmt.Sprintf(`{
		"traceId": "%s"
	}`, traceId)
}

func GetCfpCertifications_AllItem() string {
	return `[
		{
			"cfpCertificationId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"cfpCertificationDescription": "B01のCFP証明書説明。",
			"createdAt": "2024-01-01T00:00:00Z",
			"cfpCertificationFileInfo": [
				{
					"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"fileId": "fe517a2b-2af8-48ff-b1ed-88fc50f4414f",
					"fileName": "B01_CFP.pdf"
				}
			]
		}
	]`
}

func GetCfpCertifications_RequireItemOnly() string {
	return `[
		{
			"cfpCertificationId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"cfpCertificationDescription": null,
			"createdAt": "2024-01-01T00:00:00Z",
			"cfpCertificationFileInfo": [
				{
					"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"fileId": "fe517a2b-2af8-48ff-b1ed-88fc50f4414f",
					"fileName": "B01_CFP.pdf"
				}
			]
		}
	]`
}

func GetCfpCertifications_RequireItemOnlyWithUndefined() string {
	return `[
		{
			"cfpCertificationId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"createdAt": "2024-01-01T00:00:00Z",
			"cfpCertificationFileInfo": [
				{
					"operatorId": "f99c9546-e76e-9f15-35b2-abb9c9b21698",
					"fileId": "fe517a2b-2af8-48ff-b1ed-88fc50f4414f",
					"fileName": "B01_CFP.pdf"
				}
			]
		}
	]`
}

func GetCfpCertifications_NoData() string {
	return `[]`
}

func Error_PagingError() string {
	return `{
		"errors": [
			{
				"errorCode": "MSGAECO0020",
				"errorDescription": "指定した識別子は存在しません"
			}
		]
	}`
}

func Error_TraceIdNotFound() string {
	return `{
		"errors": [
			{
				"errorCode": "MSGAECI0005",
				"errorDescription": "リクエストパラメータのトレース識別子に、存在しない部品が含まれています。"
			}
		]
	}`
}

func Error_OperatorIdNotFound() string {
	return `{
		"errors": [
			{
				"errorCode": "MSGAECP0004",
				"errorDescription": "存在しない事業者識別子が使用されています。"
			}
		]
	}`
}

func Error_PartsIdNotFound() string {
	return `{
		"errors": [
			{
				"errorCode": "MSGAECP0013",
				"errorDescription": "指定された部品は存在しません。"
			}
		]
	}`
}

func Error_FileDeleteError() string {
	return `{
		"errors": [
			{
				"errorCode": "MSGAECO0034",
				"errorDescription": "ファイル削除に失敗しました。"
			}
		]
	}`
}

func Error_AuthDiffError() string {
	return `{
		"errors": [
			{
				"errorCode": "MSGAECO0025",
				"errorDescription": "認証情報と事業者識別子が一致しません。"
			}
		]
	}`
}

func Error_MaintenanceError() string {
	return `{
		"errors": [
			{
				"errorCode": "MSGXXXXYYYY",
				"errorDescription": "The service is currently undergoing maintenance. We apologize for any inconvenience."
			}
		]
	}`
}

func Error_GatewayError() string {
	return `{
		"message": "Service Unavailable"
	}`
}

func Error_BlockingPartsStructureError() string {
	return `{
		"errors": [
			{
				"errorCode": "MSGAECP0014",
				"errorDescription": "指定された部品は部品構成が存在するため削除できません。",
				"relevantData": [
					"a84012cc-73fb-4f9b-9130-59ae546f7091",
					"a84012cc-73fb-4f9b-9130-59ae546f7092"
				]
			}
		]
	}`
}

func Error_RequireError(item string) string {
	return fmt.Sprintf(`{
		"errors": [
			{
				"errorCode": "MSGAECO0001",
				"errorDescription": "%sは必須項目です。"
			}
		]
	}`, item)
}

func Error_BlockingTradeRequestError() string {
	return `{
		"errors": [
			{
				"errorCode": "MSGAECP0015",
				"errorDescription": "指定された部品は依頼済みのため削除できません。",
				"relevantData": [
					"a84012cc-73fb-4f9b-9130-59ae546f7093"
				]
			}
		]
	}`
}

func Error_BlockingTradeResponseError() string {
	return `{
		"errors": [
			{
				"errorCode": "MSGAECP0016",
				"errorDescription": "指定された部品は受領済みの依頼に紐づいているため削除できません。",
				"relevantData": [
					"a84012cc-73fb-4f9b-9130-59ae546f7094", 
					"a84012cc-73fb-4f9b-9130-59ae546f7095"
				]
			}
		]
	}`
}
