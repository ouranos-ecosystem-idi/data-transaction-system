package fixtures

import "fmt"

func GetPartsStructure_AllItem() string {
	return `{
		"parent": {
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsItem": "B01",
			"supportPartsItem": "A000001",
			"plantId": "b1234567-1234-1234-1234-123456789012",
			"operatorId": "b1234567-1234-1234-1234-123456789012",
			"amountUnitName": "kilogram",
			"endFlag": false
		},
		"children": [
			{
				"partsStructureId": "2680ed32-19a3-435b-a094-23ff43aaa611_4d987ed4-f1b0-4bf1-8795-1fdb25300e34",
				"traceId": "4d987ed4-f1b0-4bf1-8795-1fdb25300e34",
				"partsItem": "B01001",
				"supportPartsItem": "B001",
				"plantId": "b1234567-1234-1234-1234-123456789012",
				"operatorId": "b1234567-1234-1234-1234-123456789012",
				"amountUnitName": "kilogram",
				"endFlag": false,
				"amount": 2.1
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
			"plantId": "b1234567-1234-1234-1234-123456789012",
			"operatorId": "b1234567-1234-1234-1234-123456789012",
			"amountUnitName": null,
			"endFlag": false
		},
		"children": [
			{
				"partsStructureId": "2680ed32-19a3-435b-a094-23ff43aaa611_4d987ed4-f1b0-4bf1-8795-1fdb25300e34",
				"traceId": "4d987ed4-f1b0-4bf1-8795-1fdb25300e34",
				"partsItem": "B01001",
				"supportPartsItem": null,
				"plantId": "b1234567-1234-1234-1234-123456789012",
				"operatorId": "b1234567-1234-1234-1234-123456789012",
				"amountUnitName": null,
				"endFlag": false,
				"amount": null
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
			"plantId": "b1234567-1234-1234-1234-123456789012",
			"operatorId": "b1234567-1234-1234-1234-123456789012",
			"amountUnitName": null,
			"endFlag": false
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
			"plantId": "b1234567-1234-1234-1234-123456789012",
			"operatorId": "b1234567-1234-1234-1234-123456789012",
			"amountUnitName": "kg",
			"endFlag": false
		},
		"children": [
			{
				"partsStructureId": "2680ed32-19a3-435b-a094-23ff43aaa611_4d987ed4-f1b0-4bf1-8795-1fdb25300e34",
				"traceId": "4d987ed4-f1b0-4bf1-8795-1fdb25300e34",
				"partsItem": "B01001",
				"supportPartsItem": "B001",
				"plantId": "b1234567-1234-1234-1234-123456789012",
				"operatorId": "b1234567-1234-1234-1234-123456789012",
				"amountUnitName": "kg",
				"endFlag": false,
				"amount": 2.1
			}
		]
	}`
}

func PutPartsStructure() string {
	return `{
		"parent": {
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"partsItem": "B01",
			"supportPartsItem": "A000001"
		},
		"children": [
			{
				"partsStructureId": "2680ed32-19a3-435b-a094-23ff43aaa611_4d987ed4-f1b0-4bf1-8795-1fdb25300e34",
				"traceId": "4d987ed4-f1b0-4bf1-8795-1fdb25300e34",
				"partsItem": "B01001",
				"supportPartsItem": "B001"
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
				"plantId": "b1234567-1234-1234-1234-123456789012",
				"operatorId": "b1234567-1234-1234-1234-123456789012",
				"amountUnitName": "kilogram",
				"endFlag": false,
				"parentFlag": true
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
				"plantId": "b1234567-1234-1234-1234-123456789012",
				"operatorId": "b1234567-1234-1234-1234-123456789012",
				"amountUnitName": "kilogram",
				"endFlag": false,
				"parentFlag": false
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
				"plantId": "b1234567-1234-1234-1234-123456789012",
				"operatorId": "b1234567-1234-1234-1234-123456789012",
				"amountUnitName": null,
				"endFlag": false,
				"parentFlag": true
			}
		],
		"next": null
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
				"plantId": "b1234567-1234-1234-1234-123456789012",
				"operatorId": "b1234567-1234-1234-1234-123456789012",
				"amountUnitName": "kg",
				"endFlag": false,
				"parentFlag": false
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
					"requestedToOperatorId": "b1234567-1234-1234-1234-123456789012",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": "A01のCFP値を回答ください",
					"replyMessage": "A01のCFP値を回答しました"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"upstreamOperatorId": "b1234567-1234-1234-1234-123456789012",
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
					"requestedToOperatorId": "b1234567-1234-1234-1234-123456789012",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": "A01のCFP値を回答ください",
					"replyMessage": "A01のCFP値を回答しました"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"upstreamOperatorId": "b1234567-1234-1234-1234-123456789012",
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
					"requestedToOperatorId": "b1234567-1234-1234-1234-123456789012",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": null,
					"replyMessage": null
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"upstreamOperatorId": "b1234567-1234-1234-1234-123456789012",
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

func GetTradeRequests_RequireItemOnlyAnswering() string {
	return `{
		"tradeRequests": [
			{
				"request": {
					"requestId": "5185a435-c039-4196-bb34-0ee0c2395478",
					"requestType": "CFP",
					"requestStatus": "COMPLETED",
					"requestedToOperatorId": "b1234567-1234-1234-1234-123456789012",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": null,
					"replyMessage": null
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"upstreamOperatorId": "b1234567-1234-1234-1234-123456789012",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": null
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamAmountUnitName": "kilogram"
					}
				},
				"response": null
			}
		],
		"next": null
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
					"requestedToOperatorId": "b1234567-1234-1234-1234-123456789012",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": "A01のCFP値を回答ください",
					"replyMessage": "A01のCFP値を回答しました"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"downstreamOperatorId": "b1234567-1234-1234-1234-123456789012",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b"
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamPartsItem": "B01",
						"downstreamSupportPartsItem": "B0100",
						"downstreamPlantId": "b1234567-1234-1234-1234-123456789012",
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
					"requestedToOperatorId": "b1234567-1234-1234-1234-123456789012",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": "A01のCFP値を回答ください",
					"replyMessage": "A01のCFP値を回答しました"
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"downstreamOperatorId": "b1234567-1234-1234-1234-123456789012",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": "38bdd8a5-76a7-a53d-de12-725707b04a1b"
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamPartsItem": "B01",
						"downstreamSupportPartsItem": "B0100",
						"downstreamPlantId": "b1234567-1234-1234-1234-123456789012",
						"downstreamAmountUnitName": "kilogram"
					}
				}
			}
		],
		"next": null
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
					"requestedToOperatorId": "b1234567-1234-1234-1234-123456789012",
					"requestedAt": "2024-02-14T15:25:35Z",
					"requestMessage": null,
					"replyMessage": null
				},
				"trade": {
					"tradeId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
					"tradeRelation": {
						"downstreamOperatorId": "b1234567-1234-1234-1234-123456789012",
						"downstreamTraceId": "087aaa4b-8974-4a0a-9c11-b2e66ed468c5",
						"upstreamTraceId": null
					},
					"treeStatus": "UNTERMINATED",
					"downstream": {
						"downstreamPartsItem": "B01",
						"downstreamSupportPartsItem": null,
						"downstreamPlantId": "b1234567-1234-1234-1234-123456789012",
						"downstreamAmountUnitName": "kilogram"
					}
				}
			}
		],
		"next": null
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
			"replyMessage": "A01のCFP値を回答しました"
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
				}
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
				}
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

func GetCfpCertifications_AllItem() string {
	return `[
		{
			"cfpCertificationId": "a84012cc-73fb-4f9b-9130-59ae546f7092",
			"traceId": "2680ed32-19a3-435b-a094-23ff43aaa611",
			"cfpCertificationDescription": "B01のCFP証明書説明。",
			"createdAt": "2024-01-01T00:00:00Z",
			"cfpCertificationFileInfo": [
				{
					"operatorId": "b1234567-1234-1234-1234-123456789012",
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
					"operatorId": "b1234567-1234-1234-1234-123456789012",
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
