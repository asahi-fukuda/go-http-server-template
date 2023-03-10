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



/**
 * 写真アップロード用url発行要求レスポンス
 * @export
 * @interface NewUploadImageUrlsResponse
 */
export interface NewUploadImageUrlsResponse {
    /**
     * コーティングを貼り付けた状態の購入デバイス画像のアップロードurl。
     * @type {string}
     * @memberof NewUploadImageUrlsResponse
     */
    'coating_image_url'?: string;
    /**
     * コーティングを貼り付けた状態の購入デバイス画像のアップロードurlのuuid。
     * @type {string}
     * @memberof NewUploadImageUrlsResponse
     */
    'coating_image_uuid'?: string;
    /**
     * コーティング購入レシート画像のアップロードurl。
     * @type {string}
     * @memberof NewUploadImageUrlsResponse
     */
    'receipt_image_url'?: string;
    /**
     * コーティング購入レシート画像のアップロードurlのuuid。
     * @type {string}
     * @memberof NewUploadImageUrlsResponse
     */
    'receipt_image_uuid'?: string;
}

