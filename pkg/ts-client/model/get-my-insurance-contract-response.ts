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
 * 自身の保険契約情報取得レスポンス
 * @export
 * @interface GetMyInsuranceContractResponse
 */
export interface GetMyInsuranceContractResponse {
    /**
     * 
     * @type {InsuranceContract}
     * @memberof GetMyInsuranceContractResponse
     */
    'insurance_contract': InsuranceContract;
}

