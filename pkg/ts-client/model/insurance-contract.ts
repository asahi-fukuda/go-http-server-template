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


import { Contractor } from './contractor';
import { Device } from './device';

/**
 * 保険契約
 * @export
 * @interface InsuranceContract
 */
export interface InsuranceContract {
    /**
     * 個別の保険契約を識別するための id（xid）
     * @type {string}
     * @memberof InsuranceContract
     */
    'id': string;
    /**
     * 
     * @type {Device}
     * @memberof InsuranceContract
     */
    'device': Device;
    /**
     * 
     * @type {Contractor}
     * @memberof InsuranceContract
     */
    'contractor': Contractor;
    /**
     * コーティングを貼り付けた状態の購入デバイスの画像url
     * @type {string}
     * @memberof InsuranceContract
     */
    'coating_image_url'?: string;
    /**
     * 購入レシートの画像url
     * @type {string}
     * @memberof InsuranceContract
     */
    'receipt_image_url'?: string;
    /**
     * キャンセル済みの場合のみキャンセル日を表す（unixtimestamp utc）
     * @type {number}
     * @memberof InsuranceContract
     */
    'canceled_at'?: number;
}
