/* tslint:disable */
/* eslint-disable */
/**
 * Go HTTP Server Template API
 * Go HTTP Server Template API
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


// May contain unused imports in some cases
// @ts-ignore
import { Organization } from './organization';

/**
 * Organization 一覧の成功レスポンス
 * @export
 * @interface ListOrganizations200Response
 */
export interface ListOrganizations200Response {
    /**
     * Organization の配列
     * @type {Array<Organization>}
     * @memberof ListOrganizations200Response
     */
    'organizations': Array<Organization>;
}

