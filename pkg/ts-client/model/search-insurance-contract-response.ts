/* tslint:disable */
/* eslint-disable */
/**
 * HARUTO COATIONG API
 * HARUTO COATIONG API Interface
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { InsuranceContract } from './insurance-contract';

/**
 * 保険契約情報検索レスポンス
 * @export
 * @interface SearchInsuranceContractResponse
 */
export interface SearchInsuranceContractResponse {
    /**
     * 検索条件に該当する保険契約情報の配列
     * @type {Array<InsuranceContract>}
     * @memberof SearchInsuranceContractResponse
     */
    'insurance_contracts'?: Array<InsuranceContract>;
}

